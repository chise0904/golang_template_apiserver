package users

import (
	"github.com/chise0904/golang_template_apiserver/pkg/web"
)

type userSignUpRequest struct {
	Name     string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type userSignUPResponse struct {
	// UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Href     string `json:"href"`
	// Avatar   string `json:"avatar"`

}

type userSignInRequest struct {
	GrantType    string `json:"grant_type" validate:"required"`    // password, authorization_code, refresh_token
	ClientID     string `json:"client_id" validate:"required"`     // application id
	ClientSecret string `json:"client_secret" validate:"required"` // application secret
	Email        string `json:"email"`                             // 用戶 email
	Phone        string `json:"phone"`                             // 用戶 phone number
	Code         string `json:"code"`                              // ev_code pv_code
	Password     string `json:"password"`                          // 用戶 密碼
}
type userSignInResponse struct {
	UserID          string `json:"user_id"`
	Status          int32  `json:"status"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	TokenType       string `json:"token_type"`
	TokenExpiresIN  int64  `json:"token_expires_in"`
	RefreshExpireIn int64  `json:"refresh_expire_in"`
}

type setPasswordRequest struct {
	OldPassword string `json:"old_password"`                 // 用戶 舊密碼 - by password
	Password    string `json:"password" validate:"required"` // 用戶 密碼
	Phone       string `json:"phone"`                        // 用戶 phone - by verify code
	Email       string `json:"email"`                        // 用戶 email - by verify code
	Code        string `json:"code"`                         // verify code
}

type emailVerificationRequest struct {
	Email string `json:"email" validate:"required"`
	Code  string `json:"code" validate:"required"`
}
type verificationResponse struct {
	Result string `json:"result"`
}

type sendVerificationCodeRequest struct {
	Action string `json:"action" validate:"required"` // login,setpassword,revalidate,auth
	Email  string `json:"email"`                      // 用戶 email
	Phone  string `json:"phone"`                      // 用戶 phone number
}

type getAccountListRequest struct {
	By          string  `query:"by"`
	Sort        string  `query:"sort"`
	Page        uint32  `query:"page"`
	Perpage     uint32  `query:"perpage"`
	Filter      string  `query:"filter"`
	AccountType []int32 `query:"type"`
}

type getAccountListResponse struct {
	Meta     web.ResponsePayLoadMetaData `json:"meta"`
	Accounts []*AccountInfo              `json:"accounts"`
}

type updateProfileRequest struct {
	// profile
	UserName        string       `json:"user_name,omitempty"`
	Icon            []byte       `json:"icon,omitempty"`
	Description     string       `json:"description,omitempty"`
	Gender          string       `json:"gender,omitempty"`
	Birthday        *Date        `json:"birthday,omitempty"`
	Job             string       `json:"job,omitempty"`
	Country         string       `json:"country,omitempty"`
	City            string       `json:"city,omitempty"`
	District        string       `json:"district"`
	ZipCode         string       `json:"zip_code"`
	Address         string       `json:"address,omitempty"`
	ShippingAddress AddressArray `json:"shipping_address,omitempty"`
	// personal setting
	Language  string `json:"language,omitempty"`
	PhoneNoti *bool  `json:"phone_noti,omitempty"` //
	EmailNoti *bool  `json:"email_noti,omitempty"` //
	// range
	Range string `json:"range,omitempty"`
}

type changeContactsRequest struct {
	Category string `json:"category"  validate:"required"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type AccountInfo struct {
	ID              string `json:"account_id"`
	AppID           string `json:"app_id"`
	AccountType     int32  `json:"account_type"`
	AccountTypeName string `json:"account_type_name"`
	RegisMode       string `json:"regis_mode"` // default,google...etc
	Status          string `json:"status"`     //

	EvStatus   bool        `json:"ev_status"`  // email 驗證狀態
	PvStatus   bool        `json:"pv_status"`  // phone 驗證狀態
	Permission *Permission `json:"permission"` // 權限

	Password string `json:"password"` // user password
	Email    string `json:"email"`    // user email
	Phone    string `json:"phone"`    // user phone (optional)

	LoginAt   int64 `json:"login_at"`   // timestamp ms
	LogoutAt  int64 `json:"logout_at"`  // timestamp ms
	CreatedAt int64 `json:"created_at"` // timestamp ms
	UpdatedAt int64 `json:"updated_at"` // timestamp ms
	DeletedAt int64 `json:"deleted_at,omitempty"`

	//profile
	UserName string `json:"user_name"`
}

type AccountProfile struct {
	AccountID       string    `json:"account_id"`
	UserName        string    `json:"user_name"`
	Icon            []byte    `json:"icon"`
	Description     string    `json:"description"`
	Gender          string    `json:"gender"`
	Birthday        Date      `json:"birthday"`
	Job             string    `json:"job"`
	Country         string    `json:"country"`
	City            string    `json:"city"`
	District        string    `json:"district"`
	ZipCode         string    `json:"zip_code"`
	Address         string    `json:"address"`
	ShippingAddress []Address `json:"shipping_address"`

	Language  string `json:"language"`
	PhoneNoti bool   `json:"phone_noti"` //
	EmailNoti bool   `json:"email_noti"` //
	CreatedAt int64  `json:"created_at"` // timestamp ms
	UpdatedAt int64  `json:"updated_at"` // timestamp ms
	DeletedAt int64  `json:"deleted_at,omitempty"`
}
type Date struct {
	Day   int32 `json:"day,omitempty"`
	Month int32 `json:"month,omitempty"`
	Year  int32 `json:"year,omitempty"`
}

type Permission struct {
	AccessCrossAccount bool `json:"access_cross_account"`
	ProductRead        bool `json:"product_read" validate:"required"`
	ProductRewrite     bool `json:"product_rewrite" validate:"required"`
	OrderRead          bool `json:"order_read" validate:"required"`
	OrderRewrite       bool `json:"order_rewrite" validate:"required"`
	SubscribeEmail     bool `json:"subscribe_email" validate:"required"`
	CoMarketing        bool `json:"co_marketing" validate:"required"`
}

type AddressArray []Address
type Address struct {
	Type     string `json:"type"` // 宅配、超商(711.全家...etc)
	Country  string `json:"country"`
	City     string `json:"city"`
	District string `json:"district"`
	ZipCode  string `json:"zip_code"`
	Address  string `json:"address"`
	StoreID  string `json:"store_id"`
}

type listUserProfileRequest struct {
	user_ids []string `query:"user_ids"`
	Page     uint32   `query:"page"`
	Perpage  uint32   `query:"perpage"`
}

type listUserProfileResponse struct {
	Meta     web.ResponsePayLoadMetaData `json:"meta"`
	Profiles []*AccountProfile           `json:"profiles"`
}

type setAccountBlockStatusRequest struct {
	UserID    string `json:"user_id" validate:"required"`
	IsBlocked *bool  `json:"is_blocked" validate:"required"`
}
