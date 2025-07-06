package configs

import (
	"github.com/chise0904/golang_template_apiserver/pkg/web"
	"github.com/chise0904/golang_template_apiserver/pkg/zlog"

	upstream "github.com/chise0904/golang_template_apiserver/service/upstream_grpc"

	// jetstream_client "gitlab.com/hsf-cloud/lib/messaging/nats/jetstream"
	// "gitlab.com/hsf-cloud/lib/recommender/gorse"
	"go.uber.org/fx"
)

type Config struct {
	fx.Out
	// GorseConfig      *gorse.Config                `mapstructure:"gorse"`
	Log *zlog.Config `mapstructure:"log"`
	// JetstreamConfig  *jetstream_client.Config     `mapstructure:"jetstream"`
	// SSEConfig        *SSEConfig                   `mapstructure:"sse"`
	WebConfig        *web.Config                  `mapstructure:"web"`
	UpstreamGrpc     *upstream.UpstreamGrpcConfig `mapstructure:"upstream_grpc"`
	ApiGatewayConfig *ApiGatewayConfig            `mapstructure:"api_gateway"`
}

type ApiGatewayConfig struct {
	WebsiteAddr                        string          `mapstructure:"website_addr"`
	Addr                               string          `mapstructure:"addr"`
	RateLimitEnabled                   bool            `mapstructure:"rate_limit_enabled"`
	RateLimitPoolSize                  int             `mapstructure:"rate_limit_pool_size"`
	RateLimitHeaderKey                 string          `mapstructure:"rate_limit_header_key"`
	TokenGenRate                       float64         `mapstructure:"token_gen_rate"`
	TokenBucketCapacity                float64         `mapstructure:"token_bucket_capacity"`
	TokenBucketCapacityForPublicSource float64         `mapstructure:"token_bucket_capacity_for_public_source"`
	HtmlTemplatePath                   string          `mapstructure:"html_template_path"`
	ResourceConfig                     *ResourceConfig `mapstructure:"resource_config"`
}

type ResourceConfig struct {
	UsingSignURLForDownload bool  `mapstructure:"sign_url_enable"`
	SignURLDurationMin      int64 `mapstructure:"sign_url_duration_min"`
}

type SSEConfig struct {
	Key                 string `mapstructure:"host_aes_key"`
	HeartBeatIntervalMs int64  `mapstructure:"heartbeat_interval_ms"`
}
