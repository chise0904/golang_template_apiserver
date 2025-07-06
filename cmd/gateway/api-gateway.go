package apigateway

import (
	"context"
	"os"
	"time"

	"github.com/chise0904/golang_template_apiserver/configs"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/chise0904/golang_template_apiserver/delivery"

	// "gitlab.com/hsf-cloud/e-commerce/api-gateway/delivery/handler/logistics"
	// "gitlab.com/hsf-cloud/e-commerce/api-gateway/delivery/handler/order"
	// "gitlab.com/hsf-cloud/e-commerce/api-gateway/delivery/handler/product"
	// "gitlab.com/hsf-cloud/e-commerce/api-gateway/delivery/handler/sse"
	// "gitlab.com/hsf-cloud/e-commerce/api-gateway/delivery/handler/storage"
	"github.com/chise0904/golang_template_apiserver/delivery/handler/users"
	// sse_impl "gitlab.com/hsf-cloud/e-commerce/api-gateway/service/sse_service/impl"
	"github.com/chise0904/golang_template_apiserver/pkg/config"
	web "github.com/chise0904/golang_template_apiserver/pkg/web/echo"
	upstream "github.com/chise0904/golang_template_apiserver/service/upstream_grpc"

	// jetstream_client "gitlab.com/hsf-cloud/lib/messaging/nats/jetstream"
	// "gitlab.com/hsf-cloud/lib/recommender/gorse"
	"github.com/chise0904/golang_template_apiserver/pkg/zlog"
	"go.uber.org/fx"
)

func GatewayCmd() *cobra.Command {
	return &cobra.Command{
		Use: "gateway",
		Run: run,
	}
}

func run(cmd *cobra.Command, args []string) {

	cfg := configs.Config{}
	err := config.LoadConfig(os.Getenv("CONFIG_PATH"), &cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	zlog.Setup(cfg.Log)

	app := fx.New(
		fx.WithLogger(zlog.FxLogger()),
		fx.Supply(cfg),
		fx.Provide(
			// gorse.NewGorseClient,

			// upstream.SetupAIServiceGrpcClient,
			// upstream.SetupChatServiceGrpcClient,
			// upstream.SetupStorageServiceGrpcClient,
			// upstream.SetupProudctMgmtServiceGrpcClient,
			upstream.SetupUserMgmtServiceGrpcClient,
			// upstream.SetupPaymentMgmtServiceGrpcClient,
			// upstream.SetupProductRatingServiceGrpcClient,
			// upstream.SetDigitalAssetServiceGrpcClient,
			// upstream.SetupShortServiceGrpcClient,
			// upstream.SetupOrderMgmtServiceGrpcClient,
			// upstream.SetupLogisticServiceGrpcClient,
			// upstream.SetupOrderReportGrpcClient,
			// upstream.SetupUserDigitalTicketGrpcClient,

			// jetstream_client.NewJetStream,
			web.NewEcho,
			// sse_impl.NewSSE_Service,
			users.NewIdentityHandler,
			// product.NewProductHandler,
			// storage.NewStorageHandler,
			// order.NewOrderHandler,
			// sse.NewSSEHandler,
			// logistics.NewLogisticsHandler,
		),
		fx.Invoke(
			// users.LoadHtmlTemplate,
			delivery.SetDelivery,
			// sse_impl.RunServerSendEventConsumer,
		),
	)

	log.Info().Msg("launch api-gateway")
	app.Run()

	log.Info().Msg("main: shutting down api-gateway...")
	exitCode := 0
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(ctx); err != nil {
		log.Error().Msgf("main: server shutdown error: %v", err)
		exitCode++
	} else {
		log.Info().Msg("main: gracefully stopped")
	}
	os.Exit(exitCode)

}
