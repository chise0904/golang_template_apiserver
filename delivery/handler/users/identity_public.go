package users

import (
	"net/http"

	"github.com/chise0904/golang_template_apiserver/pkg/errors"
	"github.com/chise0904/golang_template_apiserver/pkg/utils"
	"github.com/chise0904/golang_template_apiserver/proto/pkg/identity"
	"github.com/labstack/echo/v4"
)

// EmailVerification implements delivery.UsersHandler.
// func (h *handler) EmailVerification(c echo.Context) error {

// 	token := c.QueryParam("access_token")

// 	r, err := h.identityGRPCClient.EmailVerification(c.Request().Context(), &identity.EmailVerificationRequest{
// 		AccessToken: token,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	htmlContent, ok := h.htmlTemplate[r.Result]
// 	if ok {
// 		return c.HTML(http.StatusOK, string(htmlContent))
// 	}

// 	// Redirect
// 	// return c.Redirect(http.StatusFound,url)
// 	return utils.MakeResponse(c, http.StatusFound, &verificationResponse{
// 		Result: r.Result.String(),
// 	})
// }

// EmailVerificationByCode implements delivery.UsersHandler.
// func (h *handler) EmailVerificationByCode(c echo.Context) error {
// 	mr := &emailVerificationRequest{}
// 	err := c.Bind(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	err = c.Validate(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	r, err := h.identityGRPCClient.EmailVerification(c.Request().Context(), &identity.EmailVerificationRequest{
// 		Email: mr.Email,
// 		Code:  mr.Code,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	htmlContent, ok := h.htmlTemplate[r.Result]
// 	if ok {
// 		return c.HTML(http.StatusOK, string(htmlContent))
// 	}

// 	return utils.MakeResponse(c, http.StatusFound, &verificationResponse{
// 		Result: r.Result.String(),
// 	})
// }

// SendVerificationCode implements delivery.UsersHandler.
func (h *handler) SendVerificationCode(c echo.Context) error {
	connection := c.QueryParam("connection")
	mr := &sendVerificationCodeRequest{}
	err := c.Bind(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}
	err = c.Validate(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}
	_, err = h.identityGRPCClient.SendVerificationCode(c.Request().Context(), &identity.SendVerificationCodeRequest{
		Connection: connection,
		Action:     mr.Action,
		Email:      mr.Email,
		Phone:      mr.Phone,
	})
	if err != nil {
		return err
	}
	return utils.MakeResponse(c, http.StatusCreated, "OK")
}

func (h *handler) SetPasswordWithoutLogin(c echo.Context) error {
	by := c.QueryParam("by")
	mr := &setPasswordRequest{}
	err := c.Bind(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}
	err = c.Validate(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}

	_, err = h.identityGRPCClient.SetPassword(c.Request().Context(), &identity.SetPasswordRequest{
		By:       by,
		Password: mr.Password,
		Phone:    mr.Phone,
		Email:    mr.Email,
		Code:     mr.Code,
	})
	if err != nil {
		return err
	}
	return utils.MakeResponse(c, http.StatusOK, "OK")
}

// UserSignIN implements delivery.UsersHandler.
func (h *handler) UserSignIN(c echo.Context) error {
	connection := c.QueryParam("connection")
	mr := &userSignInRequest{}
	err := c.Bind(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}
	err = c.Validate(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}

	r, err := h.identityGRPCClient.LoginAccount(c.Request().Context(), &identity.LoginAccountRequest{
		Connection:   connection,
		GrantType:    mr.GrantType,
		Email:        mr.Email,
		Password:     mr.Password,
		Phone:        mr.Phone,
		Code:         mr.Code,
		ClientId:     mr.ClientID,
		ClientSecret: mr.ClientSecret,
	})
	if err != nil {
		return err
	}

	if originLink := c.Request().Header.Get(constants.HttpHeader_OriginLink); originLink != "" {
		c.Response().Header().Set(constants.HttpHeader_OriginLink, originLink)
	}

	return utils.MakeResponse(c, http.StatusCreated, &userSignInResponse{
		UserID:          r.AccountId,
		Status:          int32(r.Status),
		AccessToken:     r.AccessToken,
		RefreshToken:    r.RefreshToken,
		TokenType:       r.TokenType,
		TokenExpiresIN:  r.TokenExpireIn,
		RefreshExpireIn: r.RefreshExpireIn,
	})
}

// UserSignUP implements delivery.UsersHandler.
func (h *handler) UserSignUP(c echo.Context) error {

	mr := &userSignUpRequest{}
	err := c.Bind(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}
	err = c.Validate(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}

	d, err := h.identityGRPCClient.RegisterAccount(c.Request().Context(), &identity.RegisterAccountRequest{
		Email:    mr.Email,
		Password: mr.Password,
		UserName: mr.Name,
	})
	if err != nil {
		return err
	}

	return utils.MakeResponse(c, 201, &userSignUPResponse{
		UserName: d.UserName,
		Email:    d.Email,
		Href:     d.Href,
	})
}

type refreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
type refreshAccessTokenResponse struct {
	UserID          string `json:"user_id"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	TokenType       string `json:"token_type"`
	TokenExpiresIN  int64  `json:"token_expires_in"`
	RefreshExpireIn int64  `json:"refresh_expire_in"`
}

func (h *handler) RefreshAccessToken(c echo.Context) error {

	mr := &refreshAccessTokenRequest{}
	err := c.Bind(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}
	err = c.Validate(mr)
	if err != nil {
		return errors.NewError(errors.ErrorInvalidInput, err.Error())
	}

	r, err := h.identityGRPCClient.ReFreshToken(c.Request().Context(), &identity.ReFreshTokenRequest{
		RefreshToken: mr.RefreshToken,
	})
	if err != nil {
		return err
	}

	return utils.MakeResponse(c, http.StatusCreated, &refreshAccessTokenResponse{
		UserID:          r.AccountId,
		AccessToken:     r.AccessToken,
		RefreshToken:    r.RefreshToken,
		TokenType:       r.TokenType,
		TokenExpiresIN:  r.TokenExpireIn,
		RefreshExpireIn: r.RefreshExpireIn,
	})

}
