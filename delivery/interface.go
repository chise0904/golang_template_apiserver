package delivery

import (
	"github.com/labstack/echo/v4"
	// "gitlab.com/hsf-cloud/lib/auth"
	// "gitlab.com/hsf-cloud/proto/pkg/storage_service"
)

// type StorageHelper interface {
// 	GetStorageResourceOriginalURIFunc(path string) (string, error)
// 	GetStorageResourceApiPathFunc(path string) string

// 	UploadFileHelper(ctx context.Context, fileGroup string, fileName string, fileType storage_service.StorageType, meta map[string]string, file io.ReadSeekCloser) (resp *storage_service.UploadFileStreamResponse, err error)
// 	DigitalAssetDownloader(c echo.Context, path string) error
// }

// type ProductHandler interface {
// 	CreateProduct(c echo.Context) error
// 	UpdateProduct(c echo.Context) error
// 	ListProducts(c echo.Context) error
// 	ListProductBrands(c echo.Context) error
// 	GetProduct(c echo.Context) error
// 	DeleteProduct(c echo.Context) error
// 	UpdateProductInventory(c echo.Context) error
// 	CreateProductCategory(c echo.Context) error
// 	PutProductCategory(c echo.Context) error
// 	ListProductCategories(c echo.Context) error
// 	ProductsFavoriteSetting(c echo.Context) error
// 	RecommendProducts(c echo.Context) error

// 	CreateProductRating(c echo.Context) error
// 	UpdateProductRating(c echo.Context) error
// 	ListProductRatings(c echo.Context) error
// 	ListUserProductsFavorite(c echo.Context) error

// 	ListUserDigitalProducts(c echo.Context) error
// 	ListUserDigitalAssets(c echo.Context) error
// 	GetUserDigitalAsset(c echo.Context) error
// 	GetUserDigitalAssetResource(c echo.Context) error

// 	CreateProductMaterial(c echo.Context) error
// 	ListProductMaterials(c echo.Context) error
// 	UpdateProductMaterial(c echo.Context) error
// 	DeleteProductMaterial(c echo.Context) error

// 	ListUserDigitalTickets(c echo.Context) error
// 	GetUserDigitalTicket(c echo.Context) error
// 	RedeemUserDigitalTicket(c echo.Context) error
// 	CheckCheckRedeemUserDigitalTicket(c echo.Context) error
// 	VerifyRedeemUserDigitalTicket(c echo.Context) error
// 	UserDigitalTicketHistories(c echo.Context) error
// }

// type OrderHandler interface {
// 	// shopping cart
// 	AddItemToShoppingCart(c echo.Context) error
// 	ModifyShoppingCartInfo(c echo.Context) error
// 	GetShoppingCartInfo(c echo.Context) error
// 	ClearShoppingCart(c echo.Context) error

// 	// order
// 	CallBackOrderResult(c echo.Context) error
// 	CallBackLogisticsOrderResult(c echo.Context) error
// 	CreateOrder(c echo.Context) error
// 	TradeRequestByOrderID(c echo.Context) error

// 	GetOrder(c echo.Context) error
// 	ListOrder(c echo.Context) error
// 	UpdateOrderByCustomer(c echo.Context) error
// 	UpdateOrderByManageInterface(c echo.Context) error

// 	VendorShippingSetting(c echo.Context) error
// 	ListVendorUnShippedItems(c echo.Context) error
// 	ListVendorLogisticsOrders(c echo.Context) error

// 	TradeRedirect(c echo.Context) error
// 	// order.adjustment
// 	GetAdjustmentRequest(c echo.Context) error
// 	ListAdjustmentRequest(c echo.Context) error
// 	CreateAdjustmentRequest(c echo.Context) error
// 	ReviewAdjustmentRequest(c echo.Context) error

// 	//report
// 	OrderSalesDailyReport(c echo.Context) error
// 	OrderSalesMonthlyReport(c echo.Context) error
// 	OrderSalesYearlyReport(c echo.Context) error
// 	OrderUserDailyReport(c echo.Context) error
// 	OrderUserMonthlyReport(c echo.Context) error
// 	OrderUserYearlyReport(c echo.Context) error
// 	RevenueDailyReport(c echo.Context) error
// 	RevenueMonthlyReport(c echo.Context) error
// 	RevenueYearlyReport(c echo.Context) error

// 	//For demo only
// 	DemoLogisticsNotification(c echo.Context) error
// }

// type StorageAPIGroup struct {
// 	ServerAddr       string
// 	PublicAPIGroup   *echo.Group
// 	PrivateAPIGroup  *echo.Group
// 	PublicGetAPIPath string
// }

// type StorageHandler interface {
// 	SetGetStorageResourceHandler(e *echo.Echo, pathPrefix string, g *StorageAPIGroup)
// 	SetPostStorageResourceHandler(e *echo.Echo, pathPrefix string, g *StorageAPIGroup)

// 	GetStorageResourceOriginalURIFunc(path string) (string, error)
// 	GetStorageResourceApiPathFunc(path string) string
// 	CreateShortResourceHandler(e *echo.Echo, pathPrefix string, g *StorageAPIGroup)

// 	GetShort(c echo.Context) error
// 	ListShorts(c echo.Context) error
// 	UpdateShort(c echo.Context) error
// 	DeleteShort(c echo.Context) error
// 	ShortFavoriteSetting(c echo.Context) error
// 	ListUserShortFavorite(c echo.Context) error
// 	CreateShortComment(c echo.Context) error
// 	UpdateShortComment(c echo.Context) error
// 	ListShortComments(c echo.Context) error
// 	DeleteShortComment(c echo.Context) error
// }

// type ServerSendEventHandler interface {
// 	InternalServerSendEvent(c echo.Context) error
// 	SSERegister(c echo.Context) error
// }

type UsersHandler interface {
	// UserSignUP(c echo.Context) error
	// UserSignIN(c echo.Context) error
	// EmailVerification(c echo.Context) error
	// EmailVerificationByCode(c echo.Context) error
	// SendVerificationCode(c echo.Context) error
	// SetPassword(c echo.Context) error
	// SetPasswordWithoutLogin(c echo.Context) error
	// ChangeContacts(c echo.Context) error
	// GetAccount(c echo.Context) error
	// DeleteAccount(c echo.Context) error
	// GetAccountList(c echo.Context) error
	// GetProfile(c echo.Context) error
	// ListProfiles(c echo.Context) error
	// UpdateProfile(c echo.Context) error
	// TokenVerify(c echo.Context) (*auth.UserClaims, error)
	// RefreshAccessToken(c echo.Context) error
	CreateAccount(c echo.Context) error
	// SetAccountBlockStatus(c echo.Context) error
}

// type LogisticsHandler interface {
// 	GetExpressMap(c echo.Context) error
// 	ExpressMapResult(c echo.Context) error
// 	ProductLogisticsOptions(c echo.Context) error

// 	ListLogisticsSettings(c echo.Context) error
// 	GetLogisticsSetting(c echo.Context) error

// 	PrintDeliveryNote(c echo.Context) error
// 	GetLogisticsOrder(c echo.Context) error
// }
