package upstream

import (
	"github.com/chise0904/golang_template_apiserver/pkg/grpc"
	// "gitlab.com/hsf-cloud/proto/pkg/e-commerce/chat_service"
	// "gitlab.com/hsf-cloud/proto/pkg/e-commerce/logistics_service"
	// "gitlab.com/hsf-cloud/proto/pkg/e-commerce/order_service"
	// "gitlab.com/hsf-cloud/proto/pkg/e-commerce/payment_service"
	// "gitlab.com/hsf-cloud/proto/pkg/e-commerce/product_service"
	// "gitlab.com/hsf-cloud/proto/pkg/e-commerce/short_service"
	"github.com/chise0904/golang_template_apiserver/proto/pkg/identity"
	// "gitlab.com/hsf-cloud/proto/pkg/storage_service"
)

type UpstreamGrpcConfig struct {
	// ChatServiceGrpc        string `mapstructure:"chat_service_grpc"`
	// StorageServiceGrpc     string `mapstructure:"storage_service_grpc"`
	// ProductMgmtServiceGrpc string `mapstructure:"product_mgmt_service_grpc"`
	UserMgmtServiceGrpc string `mapstructure:"user_mgmt_service_grpc"`
	// PaymentMgmtServiceGrpc string `mapstructure:"payment_mgmt_service_grpc"`
	// OrderMgmtServiceGrpc   string `mapstructure:"order_mgmt_service_grpc"`
	// LogisticServiceGrpc    string `mapstructure:"logistics_service_grpc"`
}

// func SetupAIServiceGrpcClient(cfg *UpstreamGrpcConfig) (chat_service.AIServiceClient, error) {

// 	conn, err := grpc.NewClient(cfg.ChatServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return chat_service.NewAIServiceClient(conn), nil
// }

// func SetupChatServiceGrpcClient(cfg *UpstreamGrpcConfig) (chat_service.ChatServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.ChatServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return chat_service.NewChatServiceClient(conn), nil
// }

// func SetupStorageServiceGrpcClient(cfg *UpstreamGrpcConfig) (storage_service.StorageServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.StorageServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return storage_service.NewStorageServiceClient(conn), nil
// }

// func SetupProudctMgmtServiceGrpcClient(cfg *UpstreamGrpcConfig) (product_service.ProductServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.ProductMgmtServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return product_service.NewProductServiceClient(conn), nil
// }

func SetupUserMgmtServiceGrpcClient(cfg *UpstreamGrpcConfig) (identity.IdentityServiceClient, error) {
	conn, err := grpc.NewClient(cfg.UserMgmtServiceGrpc)
	if err != nil {
		return nil, err
	}

	return identity.NewIdentityServiceClient(conn), nil
}

// func SetupPaymentMgmtServiceGrpcClient(cfg *UpstreamGrpcConfig) (payment_service.PaymentServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.PaymentMgmtServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return payment_service.NewPaymentServiceClient(conn), nil
// }

// func SetupProductRatingServiceGrpcClient(cfg *UpstreamGrpcConfig) (product_service.ProductRatingServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.ProductMgmtServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return product_service.NewProductRatingServiceClient(conn), nil
// }

// func SetDigitalAssetServiceGrpcClient(cfg *UpstreamGrpcConfig) (product_service.DigitalAssetServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.ProductMgmtServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return product_service.NewDigitalAssetServiceClient(conn), nil
// }

// func SetupShortServiceGrpcClient(cfg *UpstreamGrpcConfig) (short_service.ShortServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.StorageServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return short_service.NewShortServiceClient(conn), nil
// }

// func SetupOrderMgmtServiceGrpcClient(cfg *UpstreamGrpcConfig) (order_service.OrderServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.OrderMgmtServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return order_service.NewOrderServiceClient(conn), nil
// }

// func SetupLogisticServiceGrpcClient(cfg *UpstreamGrpcConfig) (logistics_service.LogisticServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.LogisticServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return logistics_service.NewLogisticServiceClient(conn), nil
// }

// func SetupOrderReportGrpcClient(cfg *UpstreamGrpcConfig) (order_service.OrderReportClient, error) {
// 	conn, err := grpc.NewClient(cfg.OrderMgmtServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return order_service.NewOrderReportClient(conn), nil
// }

// func SetupUserDigitalTicketGrpcClient(cfg *UpstreamGrpcConfig) (product_service.DigitalTicketServiceClient, error) {
// 	conn, err := grpc.NewClient(cfg.ProductMgmtServiceGrpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return product_service.NewDigitalTicketServiceClient(conn), nil
// }
