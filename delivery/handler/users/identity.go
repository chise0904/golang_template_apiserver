package users

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/hsf-cloud/e-commerce/api-gateway/configs"
	"gitlab.com/hsf-cloud/e-commerce/api-gateway/constants"
	"gitlab.com/hsf-cloud/e-commerce/api-gateway/delivery"
	"gitlab.com/hsf-cloud/e-commerce/api-gateway/pkg/utils"
	"gitlab.com/hsf-cloud/lib/auth"
	"gitlab.com/hsf-cloud/lib/errors"
	"gitlab.com/hsf-cloud/proto/pkg/identity"
)

type handler struct {
	apiGatewayConfig *configs.ApiGatewayConfig

	identityGRPCClient identity.IdentityServiceClient
	htmlTemplate       map[identity.VerificationResponse_VeriResult][]byte
}

func LoadHtmlTemplate(i delivery.UsersHandler, apiGatewayConfig *configs.ApiGatewayConfig) error {
	h, ok := i.(*handler)
	if !ok {
		return errors.ErrorInternalError()
	}
	if apiGatewayConfig.HtmlTemplatePath != "" {
		l := log.Logger

		err := filepath.Walk(apiGatewayConfig.HtmlTemplatePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				l.Error().Msgf("Error accessing file: %v", err)
				return err
			}

			if !info.IsDir() {
				fileName := info.Name()
				veriResult, err := strconv.ParseInt(fileName, 10, 0)
				if err != nil {
					return nil
				}

				b, err := os.ReadFile(path)
				if err != nil {
					l.Error().Msgf("read template failed: %v", err.Error())
					return err
				}
				h.htmlTemplate[identity.VerificationResponse_VeriResult(veriResult)] = b
				l.Info().Msgf("load user register %s html content", identity.VerificationResponse_VeriResult(veriResult))
			}
			return nil
		})

		if err != nil {
			return err
		}

	}

	return nil
}

func NewIdentityHandler(apiGatewayConfig *configs.ApiGatewayConfig, identityGRPCClient identity.IdentityServiceClient) delivery.UsersHandler {
	h := &handler{
		apiGatewayConfig:   apiGatewayConfig,
		identityGRPCClient: identityGRPCClient,
		htmlTemplate:       make(map[identity.VerificationResponse_VeriResult][]byte),
	}

	return h
}

func (h *handler) TokenVerify(c echo.Context) (*auth.UserClaims, error) {
	var token string
	authHeader := c.Request().Header.Get("Authorization")
	originLink := c.Request().Header.Get(constants.HttpHeader_OriginLink)
	if authHeader == "" {
		token = c.QueryParam("token")
		if token == "" {
			if originLink != "" {
				c.Response().Header().Set(constants.HttpHeader_OriginLink, originLink)
			}
			return nil, errors.ErrorNotAllow()
		}
	} else {
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return nil, errors.NewError(errors.ErrorInvalidInput, "invalid authorization header format")
		}

		token = parts[1]
	}

	resp, err := h.identityGRPCClient.CheckAccessToken(c.Request().Context(), &identity.CheckAccessTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		if originLink != "" {
			c.Response().Header().Set(constants.HttpHeader_OriginLink, originLink)
		}
		return nil, err
	}

	claims := auth.NewUserClaims(token, resp.AccountId, int32(resp.AccountType), map[string]bool{
		constants.UserClaims_PermissionKey_Can_Access_Cross_Account:     resp.Permission.CanAccessCrossAccount,
		constants.UserClaims_PermissionKey_Can_Read_Product:             resp.Permission.CanReadProduct,
		constants.UserClaims_PermissionKey_Can_Modify_Product:           resp.Permission.CanModifyProduct,
		constants.UserClaims_PermissionKey_Can_Read_Order:               resp.Permission.CanReadOrder,
		constants.UserClaims_PermissionKey_Can_Modify_Order:             resp.Permission.CanModifyOrder,
		constants.UserClaims_PermissionKey_Can_Receive_Email:            resp.Permission.CanReceiveEmails,
		constants.UserClaims_PermissionKey_Can_Participate_In_Marketing: resp.Permission.CanParticipateInMarketing,
	}, map[string]interface{}{
		constants.UserClaims_MetaKey_RegisMode: resp.RegisMode,
		constants.UserClaims_MetaKey_Status:    resp.Status,
		constants.UserClaims_MetaKey_Email:     resp.Email,
		constants.UserClaims_MetaKey_Phone:     resp.Phone,
		constants.UserClaims_MetaKey_AppID:     constants.AppID,
	})

	return claims, nil
}

type createAccountRequest struct {
	AccountType int32  `json:"account_type" validate:"required"`
	Password    string `json:"password" validate:"required"` // user password
	Email       string `json:"email" validate:"required"`    // user email
	Phone       string `json:"phone" validate:"required"`
	UserName    string `json:"user_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	EmailNoti   bool   `json:"email_noti" validate:"required"`
	PhoneNoti   bool   `json:"phone_noti" validate:"required"`
}

type createAccountResponse struct {
	AccountID   string `json:"account_id"`
	AccountType int32  `json:"account_type"`
	Email       string `json:"email"` // user email
}

func (h *handler) CreateAccount(c echo.Context) error {

	l := log.Ctx(c.Request().Context())

	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
	if err != nil {
		return err
	}
	if userClaims.GetAccountType() != int32(identity.AccountType_AccountType_ADMIN) {
		err = errors.NewError(errors.ErrorNotAllow, "Permission denied")
		return err
	}

	mr := &createAccountRequest{}
	err = c.Bind(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}
	err = c.Validate(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}

	var permission *identity.Permission
	switch mr.AccountType {
	case int32(identity.AccountType_AccountType_ADMIN):
		permission = &identity.Permission{
			CanAccessCrossAccount: true,
			CanReadProduct:        true,

			CanModifyProduct:          true,
			CanReadOrder:              true,
			CanModifyOrder:            true,
			CanReceiveEmails:          true,
			CanParticipateInMarketing: true,
		}
	case int32(identity.AccountType_AccountType_STAFF):
		permission = &identity.Permission{
			CanAccessCrossAccount:     true,
			CanReadProduct:            true,
			CanModifyProduct:          true,
			CanReadOrder:              true,
			CanModifyOrder:            true,
			CanReceiveEmails:          true,
			CanParticipateInMarketing: true,
		}
	case int32(identity.AccountType_AccountType_USER):
		permission = &identity.Permission{
			CanAccessCrossAccount:     false,
			CanReadProduct:            true,
			CanModifyProduct:          false,
			CanReadOrder:              true,
			CanModifyOrder:            true,
			CanReceiveEmails:          true,
			CanParticipateInMarketing: true,
		}
	case int32(identity.AccountType_AccountType_VENDOR):
		permission = &identity.Permission{
			CanAccessCrossAccount:     false,
			CanReadProduct:            true,
			CanModifyProduct:          false,
			CanReadOrder:              true,
			CanModifyOrder:            true,
			CanReceiveEmails:          true,
			CanParticipateInMarketing: true,
		}
	default:
		return errors.NewError(errors.ErrorInvalidInput, "unknown account type")
	}

	forLog := *mr
	forLog.Password = "********"

	l.Info().Msgf("create account: %+v", forLog)

	resp, err := h.identityGRPCClient.CreateAccount(c.Request().Context(), &identity.CreateAccountRequest{
		AppId:       constants.AppID,
		AccountType: identity.AccountType(mr.AccountType),
		RegisMode:   "NORMAL",
		Status:      identity.AccountStatus_AccountStatus_ENABLED,
		Permission:  permission,
		Password:    mr.Password,
		UserName:    mr.UserName,
		Email:       mr.Email,
		Phone:       mr.Phone,
		Gender:      mr.Gender,
		EmailNoti:   *h.convertRestfulRequestProfileBoolToProtoBoolType(&mr.EmailNoti),
		PhoneNoti:   *h.convertRestfulRequestProfileBoolToProtoBoolType(&mr.PhoneNoti),
	})

	if err != nil {
		l.Error().Msgf("CreateAccount failed: %v", err.Error())
		return err
	}

	return utils.MakeResponse(c, 201, &createAccountResponse{
		AccountType: int32(resp.AccountType),
		AccountID:   resp.AccountId,
		Email:       resp.Email,
	})
}
