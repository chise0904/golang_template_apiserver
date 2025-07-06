package delivery

import (
	"net/http/httputil"

	"github.com/chise0904/golang_template_apiserver/configs"
	"github.com/chise0904/golang_template_apiserver/pkg/utils"
	"github.com/chise0904/golang_template_apiserver/pkg/web/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	// "gitlab.com/hsf-cloud/proto/pkg/e-commerce/chat_service"
)

type handler struct {
	AppID  string
	config *configs.ApiGatewayConfig
	// chatAI           chat_service.AIServiceClient
	// chatSVC          chat_service.ChatServiceClient
	// productHandler   ProductHandler
	// storageHandler   StorageHandler
	usersHandler UsersHandler
	// orderHandler     OrderHandler
	// sseHandler       ServerSendEventHandler
	// logisticsHandler LogisticsHandler
}

func SetDelivery(
	e *echo.Echo,
	config *configs.ApiGatewayConfig,
	// chatAI chat_service.AIServiceClient,
	// chatSVC chat_service.ChatServiceClient,
	// productHandler ProductHandler,
	// storageHandler StorageHandler,
	usersHandler UsersHandler,
	// orderHandler OrderHandler,
	// sseHandler ServerSendEventHandler,
	// logisticsHandler LogisticsHandler,
) {

	h := &handler{
		config: config,
		// chatAI:           chatAI,
		// chatSVC:          chatSVC,
		// productHandler:   productHandler,
		// storageHandler:   storageHandler,
		usersHandler: usersHandler,
		// orderHandler:     orderHandler,
		// sseHandler:       sseHandler,
		// logisticsHandler: logisticsHandler,
	}

	setRouter(config, e, h)

}

func setRouter(config *configs.ApiGatewayConfig, e *echo.Echo, h *handler) {

	var apiRatelimitMiddle echo.MiddlewareFunc
	// var sourceApiRatelimitMiddle echo.MiddlewareFunc

	if config.RateLimitEnabled {
		if config.RateLimitHeaderKey == "" {
			log.Info().Msgf("api rate limit enabled,by real ip rate: %v, capacity: %v, capacity for public source: %v", config.TokenGenRate, config.TokenBucketCapacity, config.TokenBucketCapacityForPublicSource)
		} else {
			log.Info().Msgf("api rate limit enabled,by header key: %s rate: %v, capacity: %v, capacity for public source: %v", config.RateLimitHeaderKey, config.TokenGenRate, config.TokenBucketCapacity, config.TokenBucketCapacityForPublicSource)
		}
		apiRatelimitMiddle = utils.RateLimitByClientIpMiddleware(config.RateLimitPoolSize, config.TokenGenRate, config.TokenBucketCapacity, config.RateLimitHeaderKey)
		// sourceApiRatelimitMiddle = utils.RateLimitByClientIpMiddleware(config.RateLimitPoolSize, config.TokenGenRate, config.TokenBucketCapacityForPublicSource, config.RateLimitHeaderKey)

	}

	apiV1PublicGroup := e.Group("/public/apis/v1", apiRatelimitMiddle)
	apiV1PrivateGroup := e.Group("/apis/v1", apiRatelimitMiddle, middleware.JWTAuthMiddlewareFunc(h.usersHandler)) // access token needed
	// apiV1InternalGroup := e.Group("/internal/apis/v1")                                                             // internal service access only

	// apiV1PublicSourceGroup := e.Group("/public/source", sourceApiRatelimitMiddle)

	apiV1PublicGroup.Any("/echo/*", func(c echo.Context) error {
		dumpReq, _ := httputil.DumpRequest(c.Request(), true)

		l := log.Ctx(c.Request().Context())
		l.Info().Msgf("request dump: %s", dumpReq)

		return nil
	})

	// //order
	// // shopping cart
	// apiV1PrivateGroup.Add("GET", "/cart", h.orderHandler.GetShoppingCartInfo)
	// // apiV1PrivateGroup.Add("POST", "/cart", h.orderHandler.AddItemToShoppingCart)
	// apiV1PrivateGroup.Add("PUT", "/cart", h.orderHandler.ModifyShoppingCartInfo)
	// apiV1PrivateGroup.Add("DELETE", "/cart", h.orderHandler.ClearShoppingCart)

	// // order trade call back
	// apiV1PublicGroup.Add("POST", "/orders/trade/callback", h.orderHandler.CallBackOrderResult)
	// apiV1PublicGroup.Add("GET", "/orders/trade_redirect", h.orderHandler.TradeRedirect)
	// apiV1PublicGroup.Add("POST", "/orders/logistics/callback", h.orderHandler.CallBackLogisticsOrderResult)

	// apiV1PrivateGroup.Add("POST", "/orders", h.orderHandler.CreateOrder)
	// apiV1PrivateGroup.Add("GET", "/orders/:orderNo", h.orderHandler.GetOrder)
	// apiV1PrivateGroup.Add("GET", "/orders", h.orderHandler.ListOrder)
	// apiV1PrivateGroup.Add("POST", "/orders/:orderNo", h.orderHandler.UpdateOrderByCustomer)

	// apiV1PrivateGroup.Add("POST", "/orders/vendor/setting", h.orderHandler.VendorShippingSetting)
	// apiV1PrivateGroup.Add("GET", "/orders/vendor/items", h.orderHandler.ListVendorUnShippedItems)
	// apiV1PrivateGroup.Add("GET", "/orders/vendor/logistics", h.orderHandler.ListVendorLogisticsOrders)

	// apiV1PrivateGroup.Add("POST", "/orders/manage", h.orderHandler.UpdateOrderByManageInterface)
	// // order adjustment request
	// apiV1PrivateGroup.Add("POST", "/orders/adjustment/manage", h.orderHandler.ReviewAdjustmentRequest)

	// apiV1PrivateGroup.Add("POST", "/orders/adjustment", h.orderHandler.CreateAdjustmentRequest)
	// apiV1PrivateGroup.Add("GET", "/orders/adjustment/:id", h.orderHandler.GetAdjustmentRequest)
	// apiV1PrivateGroup.Add("GET", "/orders/adjustment", h.orderHandler.ListAdjustmentRequest)

	// //report
	// apiV1PrivateGroup.Add("GET", "/orders/reports/sales/:year", h.orderHandler.OrderSalesYearlyReport)
	// apiV1PrivateGroup.Add("GET", "/orders/reports/sales/:year/:month", h.orderHandler.OrderSalesMonthlyReport)
	// apiV1PrivateGroup.Add("GET", "/orders/reports/sales/:year/:month/:day", h.orderHandler.OrderSalesDailyReport)
	// apiV1PrivateGroup.Add("GET", "/orders/reports/users/:year", h.orderHandler.OrderUserYearlyReport)
	// apiV1PrivateGroup.Add("GET", "/orders/reports/users/:year/:month", h.orderHandler.OrderUserMonthlyReport)
	// apiV1PrivateGroup.Add("GET", "/orders/reports/users/:year/:month/:day", h.orderHandler.OrderUserDailyReport)
	// apiV1PrivateGroup.Add("GET", "/orders/reports/revenue/:year", h.orderHandler.RevenueYearlyReport)
	// apiV1PrivateGroup.Add("GET", "/orders/reports/revenue/:year/:month", h.orderHandler.RevenueMonthlyReport)
	// apiV1PrivateGroup.Add("GET", "/orders/reports/revenue/:year/:month/:day", h.orderHandler.RevenueDailyReport)

	// users
	apiV1PublicGroup.Add("POST", "/users/accounts/signUp", h.usersHandler.UserSignUP)
	// apiV1PublicGroup.Add("POST", "/users/accounts/passwords", h.usersHandler.SetPasswordWithoutLogin)

	// apiV1PublicGroup.Add("GET", "/users/accounts/emailVerification", h.usersHandler.EmailVerification)
	// apiV1PublicGroup.Add("POST", "/users/accounts/emailVerification", h.usersHandler.EmailVerificationByCode)
	apiV1PublicGroup.Add("POST", "/users/oauth2/signIn", h.usersHandler.UserSignIN)
	// apiV1PublicGroup.Add("POST", "/users/oauth2/refresh", h.usersHandler.RefreshAccessToken)
	// apiV1PublicGroup.Add("POST", "/users/oauth2/verificationCodes", h.usersHandler.SendVerificationCode)

	// apiV1PrivateGroup.Add("POST", "/users/accounts/passwords", h.usersHandler.SetPassword)
	// apiV1PrivateGroup.Add("GET", "/users/accounts", h.usersHandler.GetAccountList)
	// apiV1PrivateGroup.Add("GET", "/users/accounts/:accountID", h.usersHandler.GetAccount)
	// apiV1PrivateGroup.Add("DELETE", "/users/accounts/:accountID", h.usersHandler.DeleteAccount)
	// apiV1PrivateGroup.Add("GET", "/users/accounts/profiles/:accountID", h.usersHandler.GetProfile)
	// apiV1PrivateGroup.Add("GET", "/users/accounts/profiles", h.usersHandler.ListProfiles)
	// apiV1PrivateGroup.Add("PUT", "/users/accounts/profiles/:accountID", h.usersHandler.UpdateProfile)
	// apiV1PrivateGroup.Add("POST", "/users/accounts/:accountID/contacts", h.usersHandler.ChangeContacts)
	apiV1PrivateGroup.Add("POST", "/users/accounts", h.usersHandler.CreateAccount)

	// apiV1PrivateGroup.Add("POST", "/users/accounts/manage", h.usersHandler.SetAccountBlockStatus)

	// apiV1InternalGroup.Add("GET", "/users/oauth2/userInfo", h.usersHandler.CheckAccessToken)

	//product
	// apiV1PrivateGroup.Add("POST", "/products", h.productHandler.CreateProduct)
	// apiV1PrivateGroup.Add("PUT", "/products/:id", h.productHandler.UpdateProduct)
	// apiV1PrivateGroup.Add("PUT", "/products/:id/inventory", h.productHandler.UpdateProductInventory)
	// apiV1PrivateGroup.Add("GET", "/products", h.productHandler.ListProducts)
	// apiV1PrivateGroup.Add("GET", "/products/:id", h.productHandler.GetProduct)
	// apiV1PrivateGroup.Add("DELETE", "/products/:id", h.productHandler.DeleteProduct)
	// apiV1PrivateGroup.Add("POST", "/products/favorite", h.productHandler.ProductsFavoriteSetting)
	// apiV1PrivateGroup.Add("GET", "/products/favorite", h.productHandler.ListUserProductsFavorite)
	// apiV1PrivateGroup.Add("GET", "/products/brands", h.productHandler.ListProductBrands)
	// apiV1PrivateGroup.Add("POST", "/products/:id/ratings", h.productHandler.CreateProductRating)
	// apiV1PrivateGroup.Add("PUT", "/products/ratings", h.productHandler.UpdateProductRating)
	// apiV1PublicGroup.Add("GET", "/products/:id/ratings", h.productHandler.ListProductRatings)
	// apiV1PrivateGroup.Add("GET", "/products/:id/ratings", h.productHandler.ListProductRatings)

	// apiV1PrivateGroup.Add("GET", "/products/recommend_products", h.productHandler.RecommendProducts)
	// apiV1PublicGroup.Add("GET", "/products/recommend_products", h.productHandler.RecommendProducts)
	// apiV1PrivateGroup.Add("GET", "/products/logistics_options", h.logisticsHandler.ProductLogisticsOptions)

	// //user digital asset
	// apiV1PrivateGroup.Add("GET", "/assets", h.productHandler.ListUserDigitalAssets)
	// apiV1PrivateGroup.Add("GET", "/products/assets", h.productHandler.ListUserDigitalProducts)

	// apiV1PrivateGroup.Add("GET", "/assets/:id", h.productHandler.GetUserDigitalAsset)
	// apiV1PrivateGroup.Add("GET", "/assets/:id/source/:meta", h.productHandler.GetUserDigitalAssetResource)

	// //user digital ticket
	// apiV1PrivateGroup.Add("GET", "/tickets", h.productHandler.ListUserDigitalTickets)
	// apiV1PrivateGroup.Add("GET", "/tickets/histories", h.productHandler.UserDigitalTicketHistories)
	// apiV1PrivateGroup.Add("GET", "/tickets/:id", h.productHandler.GetUserDigitalTicket)
	// apiV1PrivateGroup.Add("POST", "/tickets/:id/redeem", h.productHandler.RedeemUserDigitalTicket)
	// apiV1PrivateGroup.Add("GET", "/tickets/:id/check", h.productHandler.CheckCheckRedeemUserDigitalTicket)
	// apiV1PrivateGroup.Add("POST", "/tickets/:id/verify", h.productHandler.VerifyRedeemUserDigitalTicket)
	// apiV1PrivateGroup.Add("GET", "/tickets/:id/histories", h.productHandler.UserDigitalTicketHistories)

	// //product category
	// apiV1PrivateGroup.Add("POST", "/products/categories", h.productHandler.CreateProductCategory)
	// apiV1PrivateGroup.Add("PUT", "/products/categories", h.productHandler.PutProductCategory)
	// apiV1PrivateGroup.Add("GET", "/products/categories", h.productHandler.ListProductCategories)

	// //product public api
	// apiV1PublicGroup.Add("GET", "/products", h.productHandler.ListProducts)
	// apiV1PublicGroup.Add("GET", "/products/:id", h.productHandler.GetProduct)
	// apiV1PublicGroup.Add("GET", "/products/brands", h.productHandler.ListProductBrands)
	// apiV1PublicGroup.Add("GET", "/products/categories", h.productHandler.ListProductCategories)

	// apiV1PrivateGroup.Add("GET", "/products/:id/materials", h.productHandler.ListProductMaterials)
	// apiV1PrivateGroup.Add("POST", "/products/:id/materials", h.productHandler.CreateProductMaterial)
	// apiV1PrivateGroup.Add("PUT", "/products/:id/materials/:material_id", h.productHandler.UpdateProductMaterial)
	// apiV1PrivateGroup.Add("DELETE", "/products/:id/materials/:material_id", h.productHandler.DeleteProduct)

	//storage
	// storageAPI := &StorageAPIGroup{
	// 	ServerAddr:      h.config.Addr,
	// 	PublicAPIGroup:  apiV1PublicSourceGroup,
	// 	PrivateAPIGroup: apiV1PrivateGroup,
	// }
	// h.storageHandler.SetGetStorageResourceHandler(e, "/storage", storageAPI)
	// h.storageHandler.SetPostStorageResourceHandler(e, "/storage", storageAPI)

	//shorts
	// h.storageHandler.CreateShortResourceHandler(e, "/shorts", storageAPI)
	// apiV1PublicGroup.GET("/shorts", h.storageHandler.ListShorts)
	// apiV1PublicGroup.GET("/shorts/:id", h.storageHandler.GetShort)
	// apiV1PrivateGroup.GET("/shorts", h.storageHandler.ListShorts)
	// apiV1PrivateGroup.GET("/shorts/:id", h.storageHandler.GetShort)
	// apiV1PrivateGroup.POST("/shorts/:id", h.storageHandler.UpdateShort)
	// apiV1PrivateGroup.DELETE("/shorts/:id", h.storageHandler.DeleteShort)

	// apiV1PrivateGroup.POST("/shorts/favorite", h.storageHandler.ShortFavoriteSetting)
	// apiV1PrivateGroup.GET("/shorts/favorite", h.storageHandler.ListUserShortFavorite)
	// apiV1PrivateGroup.POST("/shorts/:id/comments", h.storageHandler.CreateShortComment)
	// apiV1PrivateGroup.PUT("/shorts/comments", h.storageHandler.UpdateShortComment)
	// apiV1PrivateGroup.GET("/shorts/:id/comments", h.storageHandler.ListShortComments)
	// apiV1PrivateGroup.DELETE("/shorts/comments", h.storageHandler.DeleteShortComment)

	//SSE
	// apiV1PrivateGroup.Any("/events", h.sseHandler.SSERegister)
	// apiV1InternalGroup.POST("/events/send", h.sseHandler.InternalServerSendEvent)

	// //Logistics
	// apiV1PrivateGroup.GET("/logistics/express_map", h.logisticsHandler.GetExpressMap)
	// apiV1PublicGroup.Any("/logistics/callback/map", h.logisticsHandler.ExpressMapResult)
	// apiV1PrivateGroup.GET("/logistics/orders/:logistics_order_id", h.logisticsHandler.GetLogisticsOrder)
	// apiV1PrivateGroup.GET("/logistics/delivery_note/:logistics_order_id", h.logisticsHandler.PrintDeliveryNote)
	// apiV1PrivateGroup.GET("/logistics/settings/:setting_type", h.logisticsHandler.ListLogisticsSettings)
	// apiV1PrivateGroup.GET("/logistics/settings/:setting_type/:logistics_type/:logistics_sub_type", h.logisticsHandler.GetLogisticsSetting)

	//for test or demo...
	// apiV1PrivateGroup.POST("/orders/logistics/notification", h.orderHandler.DemoLogisticsNotification)

	rs := e.Routes()
	for i := 0; i < len(rs); i++ {
		v := rs[i]
		log.Debug().Msgf("(%d) [%s] %s - %s", i+1, v.Method, v.Path, v.Name)
	}

}
