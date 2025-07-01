package users

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/labstack/echo/v4"
// 	"gitlab.com/hsf-cloud/e-commerce/api-gateway/constants"
// 	"gitlab.com/hsf-cloud/e-commerce/api-gateway/pkg/utils"
// 	"gitlab.com/hsf-cloud/lib/auth"
// 	"gitlab.com/hsf-cloud/lib/errors"
// 	"gitlab.com/hsf-cloud/lib/pagination"
// 	"gitlab.com/hsf-cloud/lib/web"
// 	"gitlab.com/hsf-cloud/proto/pkg/common"
// 	"gitlab.com/hsf-cloud/proto/pkg/identity"
// )

// // =================== private ===================
// // ChangeContacts implements delivery.UsersHandler. (private)
// func (h *handler) ChangeContacts(c echo.Context) error {

// 	mr := &changeContactsRequest{}
// 	err := c.Bind(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	err = c.Validate(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	id := c.Param("accountID")

// 	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
// 	if err != nil {
// 		return err
// 	}
// 	if userClaims.GetUserID() != id {
// 		if userClaims.GetAccountType() != int32(identity.AccountType_AccountType_ADMIN) {
// 			err = errors.NewError(errors.ErrorNotAllow, "Permission denied")
// 			return err
// 		}
// 	}

// 	_, err = h.identityGRPCClient.ChangeContacts(c.Request().Context(), &identity.ChangeContactsRequest{
// 		AccountId: id,
// 		Email:     mr.Email,
// 		Phone:     mr.Phone,
// 		Category:  mr.Category,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return utils.MakeResponse(c, http.StatusOK, "OK")
// }

// // DeleteAccount implements delivery.UsersHandler. (private)
// func (h *handler) DeleteAccount(c echo.Context) error {
// 	id := c.Param("accountID")

// 	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
// 	if err != nil {
// 		return err
// 	}
// 	if userClaims.GetUserID() != id {
// 		if userClaims.GetAccountType() != int32(identity.AccountType_AccountType_ADMIN) {
// 			err = errors.NewError(errors.ErrorNotAllow, "Permission denied")
// 			return err
// 		}
// 	}

// 	_, err = h.identityGRPCClient.DeleteOneAccount(c.Request().Context(), &identity.DeleteOneAccountRequest{
// 		AccountId: id,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return utils.MakeResponse(c, http.StatusOK, "OK")
// }

// // GetAccount implements delivery.UsersHandler. (private)
// func (h *handler) GetAccount(c echo.Context) error {
// 	id := c.Param("accountID")

// 	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
// 	if err != nil {
// 		return err
// 	}
// 	if userClaims.GetUserID() != id {
// 		if userClaims.GetAccountType() != int32(identity.AccountType_AccountType_ADMIN) {
// 			err = errors.NewError(errors.ErrorNotAllow, "Permission denied")
// 			return err
// 		}
// 	}
// 	r, err := h.identityGRPCClient.GetOneAccount(c.Request().Context(), &identity.GetOneAccountRequest{
// 		AccountId: id,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	resp := h.convertProtoAccountInfoValueToJson(userClaims, r)

// 	return utils.MakeResponse(c, http.StatusOK, resp)
// }

// // GetAccountList implements delivery.UsersHandler. (private)
// func (h *handler) GetAccountList(c echo.Context) error {
// 	mr := &getAccountListRequest{}
// 	err := c.Bind(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	err = c.Validate(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
// 	if err != nil {
// 		return err
// 	}

// 	var accountTypes []identity.AccountType
// 	for _, v := range mr.AccountType {
// 		accountTypes = append(accountTypes, identity.AccountType(v))
// 	}

// 	filter := &identity.GetAllAccountRequest{
// 		By:           mr.By,
// 		Sort:         mr.Sort,
// 		Page:         mr.Page,
// 		Perpage:      mr.Perpage,
// 		Filter:       mr.Filter,
// 		AccountTypes: accountTypes,
// 	}

// 	if !userClaims.GetPermission()[constants.UserClaims_PermissionKey_Can_Access_Cross_Account] {
// 		filter.UserIds = []string{userClaims.GetUserID()}
// 	}

// 	acs, err := h.identityGRPCClient.GetAllAccount(c.Request().Context(), filter)
// 	if err != nil {
// 		return err
// 	}

// 	resp := &getAccountListResponse{
// 		Meta: web.ResponsePayLoadMetaData{
// 			Pagination: &pagination.Pagination{
// 				Page:       uint32(acs.Meta.Page),
// 				PerPage:    uint32(acs.Meta.Perpage),
// 				TotalCount: uint32(acs.Meta.TotalCount),
// 				TotalPage:  uint32(acs.Meta.TotalPage),
// 			},
// 		},
// 	}

// 	var accountIDs []string
// 	for _, v := range acs.Accounts {
// 		accountIDs = append(accountIDs, v.AccountId)
// 		resp.Accounts = append(resp.Accounts, h.convertProtoAccountInfoValueToJson(userClaims, v))
// 	}

// 	if len(accountIDs) > 0 {
// 		profiles, err := h.identityGRPCClient.ListProfiles(c.Request().Context(), &identity.ListProfilesRequest{
// 			AccountIds: accountIDs,
// 			Page:       1,
// 			Perpage:    uint32(len(accountIDs)),
// 		})
// 		if err == nil {
// 			for _, p := range profiles.Profiles {
// 				for _, a := range resp.Accounts {
// 					if a.ID == p.AccountId {
// 						a.UserName = p.UserName
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return utils.MakeResponse(c, http.StatusOK, resp)

// }

// // GetProfile implements delivery.UsersHandler. (private)
// func (h *handler) GetProfile(c echo.Context) error {
// 	id := c.Param("accountID")

// 	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
// 	if err != nil {
// 		return err
// 	}
// 	if userClaims.GetUserID() != id {
// 		if userClaims.GetAccountType() != int32(identity.AccountType_AccountType_ADMIN) {
// 			err = errors.NewError(errors.ErrorNotAllow, "Permission denied")
// 			return err
// 		}
// 	}

// 	r, err := h.identityGRPCClient.GetProfile(c.Request().Context(), &identity.GetProfileRequest{
// 		AccountId: id,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	resp := h.convertProtoProfileValueToJson(r)

// 	return utils.MakeResponse(c, http.StatusOK, resp)
// }

// // SetPassword implements delivery.UsersHandler. (private)
// func (h *handler) SetPassword(c echo.Context) error {
// 	by := c.QueryParam("by")
// 	mr := &setPasswordRequest{}
// 	err := c.Bind(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	err = c.Validate(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
// 	if err != nil {
// 		return err
// 	}

// 	_, err = h.identityGRPCClient.SetPassword(c.Request().Context(), &identity.SetPasswordRequest{
// 		By:          by,
// 		AccessToken: userClaims.GetAccessToken(),
// 		Password:    mr.Password,
// 		OldPassword: mr.OldPassword,
// 		Phone:       mr.Phone,
// 		Email:       mr.Email,
// 		Code:        mr.Code,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return utils.MakeResponse(c, http.StatusOK, "OK")
// }

// // UpdateProfile implements delivery.UsersHandler. (private)
// func (h *handler) UpdateProfile(c echo.Context) error {

// 	mr := &updateProfileRequest{}
// 	err := c.Bind(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	err = c.Validate(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	id := c.Param("accountID")
// 	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
// 	if err != nil {
// 		return err
// 	}
// 	if userClaims.GetUserID() != id {
// 		if userClaims.GetAccountType() != int32(identity.AccountType_AccountType_ADMIN) {
// 			err = errors.NewError(errors.ErrorNotAllow, "Permission denied")
// 			return err
// 		}
// 	}

// 	_, err = h.identityGRPCClient.UpdateProfile(c.Request().Context(), &identity.UpdateProfileRequest{
// 		AccountId:       id,
// 		UserName:        mr.UserName,
// 		Icon:            mr.Icon,
// 		Description:     mr.Description,
// 		Gender:          mr.Gender,
// 		Birthday:        h.convertRestfulRequestBirthdayValueToProtoBirthday(mr.Birthday),
// 		Job:             mr.Job,
// 		Country:         mr.Country,
// 		City:            mr.City,
// 		District:        mr.District,
// 		ZipCode:         mr.ZipCode,
// 		Address:         mr.Address,
// 		ShippingAddress: h.convertRestfulRequestShippingAddressValueToProtoAddressArray(mr.ShippingAddress),
// 		Language:        mr.Language,
// 		EmailNoti:       *h.convertRestfulRequestProfileBoolToProtoBoolType(mr.EmailNoti),
// 		PhoneNoti:       *h.convertRestfulRequestProfileBoolToProtoBoolType(mr.PhoneNoti),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return utils.MakeResponse(c, http.StatusOK, "OK")
// }

// // SetAccountBlockStatus implements delivery.UsersHandler. (private)
// func (h *handler) SetAccountBlockStatus(c echo.Context) error {

// 	mr := &setAccountBlockStatusRequest{}
// 	err := c.Bind(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	err = c.Validate(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	userClaims, err := auth.GetUserClaimsForContext(c.Request().Context())
// 	if err != nil {
// 		return err
// 	}

// 	if userClaims.GetAccountType() != int32(identity.AccountType_AccountType_ADMIN) {
// 		err = errors.NewError(errors.ErrorNotAllow, "Permission denied")
// 		return err
// 	}

// 	_, err = h.identityGRPCClient.SetAccountBlockStatus(c.Request().Context(), &identity.SetAccountBlockStatusRequest{
// 		AccountId: mr.UserID,
// 		IsBlocked: *mr.IsBlocked,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return utils.MakeResponse(c, http.StatusOK, "OK")
// }

// func (h *handler) convertRestfulRequestShippingAddressValueToProtoAddressArray(r []Address) []*identity.Adderss {
// 	var result []*identity.Adderss
// 	for _, v := range r {

// 		s := &identity.Adderss{
// 			Type:     v.Type,
// 			Country:  v.Country,
// 			City:     v.City,
// 			District: v.District,
// 			ZipCode:  v.ZipCode,
// 			Address:  v.Address,
// 			StoreId:  v.StoreID,
// 		}

// 		result = append(result, s)
// 	}
// 	return result
// }
// func (h *handler) convertProtoAddressArrayValueToJsonShippingAddress(r []*identity.Adderss) []Address {
// 	var result []Address
// 	for _, v := range r {

// 		s := Address{
// 			Type:     v.Type,
// 			Country:  v.Country,
// 			City:     v.City,
// 			District: v.District,
// 			ZipCode:  v.ZipCode,
// 			Address:  v.Address,
// 			StoreID:  v.StoreId,
// 		}

// 		result = append(result, s)
// 	}
// 	return result
// }
// func (h *handler) convertRestfulRequestBirthdayValueToProtoBirthday(r *Date) *identity.Date {
// 	var result *identity.Date
// 	if r != nil {
// 		result = &identity.Date{
// 			Day:   r.Day,
// 			Month: r.Month,
// 			Year:  r.Year,
// 		}
// 	} else {
// 		return nil
// 	}
// 	return result
// }
// func (h *handler) convertProtoBirthdayValueToJsonBirthday(r *identity.Date) *Date {
// 	var result *Date
// 	if r != nil {
// 		result = &Date{
// 			Day:   r.Day,
// 			Month: r.Month,
// 			Year:  r.Year,
// 		}
// 	} else {
// 		return nil
// 	}
// 	return result
// }

// func (h *handler) convertRestfulRequestProfileBoolToProtoBoolType(r *bool) *common.BoolType {
// 	var result *common.BoolType

// 	if r == nil {
// 		result = common.BoolType_NoSet.Enum()
// 	} else if *r {
// 		result = common.BoolType_True.Enum()
// 	} else if !*r {
// 		result = common.BoolType_False.Enum()
// 	}
// 	return result
// }

// func (h *handler) convertProtoPermissionValueToJson(r *identity.Permission) *Permission {
// 	return &Permission{
// 		AccessCrossAccount: r.CanAccessCrossAccount,
// 		ProductRead:        r.CanReadProduct,
// 		ProductRewrite:     r.CanModifyProduct,
// 		OrderRead:          r.CanReadOrder,
// 		OrderRewrite:       r.CanModifyOrder,
// 		SubscribeEmail:     r.CanReceiveEmails,
// 		CoMarketing:        r.CanParticipateInMarketing,
// 	}
// }

// func (h *handler) convertProtoAccountInfoValueToJson(claims *auth.UserClaims, r *identity.AccountInfo) *AccountInfo {

// 	password := "*********"
// 	// if claims != nil && claims.GetAccountType() == int32(identity.AccountType_AccountType_ADMIN) {
// 	// 	password = r.Password
// 	// }
// 	result := &AccountInfo{
// 		ID:              r.AccountId,
// 		AppID:           r.AppId,
// 		AccountType:     int32(r.AccountType),
// 		AccountTypeName: strings.ReplaceAll(r.AccountType.String(), "AccountType_", ""),
// 		RegisMode:       r.RegisMode,
// 		Status:          r.Status.String(),
// 		EvStatus:        r.EvStatus,
// 		PvStatus:        r.PvStatus,
// 		Permission:      h.convertProtoPermissionValueToJson(r.Permission),
// 		Password:        password,
// 		Email:           r.Email,
// 		Phone:           r.Phone,
// 		LoginAt:         r.LoginAt,
// 		LogoutAt:        r.LogoutAt,
// 		CreatedAt:       r.CreatedAt,
// 		UpdatedAt:       r.UpdatedAt,
// 		DeletedAt:       r.DeletedAt,
// 	}

// 	return result
// }

// func (h *handler) convertProtoProfileValueToJson(r *identity.UserProfile) *AccountProfile {

// 	result := &AccountProfile{
// 		AccountID:       r.AccountId,
// 		UserName:        r.UserName,
// 		Icon:            r.Icon,
// 		Description:     r.Description,
// 		Gender:          r.Gender,
// 		Birthday:        *h.convertProtoBirthdayValueToJsonBirthday(r.Birthday),
// 		Job:             r.Job,
// 		Country:         r.Country,
// 		City:            r.City,
// 		District:        r.District,
// 		ZipCode:         r.ZipCode,
// 		Address:         r.Address,
// 		ShippingAddress: h.convertProtoAddressArrayValueToJsonShippingAddress(r.ShippingAddress),
// 		Language:        r.Language,
// 		PhoneNoti:       r.PhoneNoti,
// 		EmailNoti:       r.EmailNoti,
// 		CreatedAt:       r.CreatedAt,
// 		UpdatedAt:       r.UpdatedAt,
// 		DeletedAt:       r.DeletedAt,
// 	}

// 	return result
// }

// func (h *handler) ListProfiles(c echo.Context) error {

// 	mr := &listUserProfileRequest{}
// 	err := c.Bind(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	err = c.Validate(mr)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	result, err := h.identityGRPCClient.ListProfiles(c.Request().Context(), &identity.ListProfilesRequest{
// 		AccountIds: mr.user_ids,
// 		Page:       mr.Page,
// 		Perpage:    mr.Perpage,
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	var out []*AccountProfile
// 	for _, v := range result.Profiles {
// 		out = append(out, &AccountProfile{
// 			AccountID:       v.AccountId,
// 			UserName:        v.UserName,
// 			Icon:            v.Icon,
// 			Description:     v.Description,
// 			Gender:          v.Gender,
// 			Birthday:        *h.convertProtoBirthdayValueToJsonBirthday(v.Birthday),
// 			Job:             v.Job,
// 			Country:         v.Country,
// 			City:            v.City,
// 			District:        v.District,
// 			ZipCode:         v.ZipCode,
// 			Address:         v.Address,
// 			ShippingAddress: h.convertProtoAddressArrayValueToJsonShippingAddress(v.ShippingAddress),
// 			Language:        v.Language,
// 			PhoneNoti:       v.PhoneNoti,
// 			EmailNoti:       v.EmailNoti,
// 			CreatedAt:       v.CreatedAt,
// 			UpdatedAt:       v.UpdatedAt,
// 			DeletedAt:       v.DeletedAt,
// 		})
// 	}

// 	resp := &listUserProfileResponse{
// 		Meta: web.ResponsePayLoadMetaData{
// 			Pagination: &pagination.Pagination{
// 				TotalCount: uint32(result.Meta.TotalCount),
// 				TotalPage:  uint32(result.Meta.TotalPage),
// 				Page:       uint32(result.Meta.Page),
// 				PerPage:    uint32(result.Meta.Perpage),
// 			},
// 		},
// 		Profiles: out,
// 	}
// 	return utils.MakeResponse(c, http.StatusOK, resp)
// }
