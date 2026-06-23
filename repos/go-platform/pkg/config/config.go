package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

type GatewayConfig struct {
	AppName      string     `env:"APP_NAME" envDefault:"gateway"`
	GoEnv        string     `env:"GO_ENV" envDefault:"development" validate:"oneof=development production"`
	Port         string     `env:"GATEWAY_PORT" envDefault:"8080"`
	MetricsPort  string     `env:"GATEWAY_METRICS_PORT" envDefault:"9090"`
	UserGRPCHost string     `env:"USER_GRPC_HOST" envDefault:"localhost"`
	UserGRPCPort string     `env:"USER_GRPC_PORT" envDefault:"50051"`
	Cors         CorsConfig `envPrefix:"CORS_"`
	Otel         OtelConfig `envPrefix:"OTEL_"`
	Log          LogConfig  `envPrefix:"LOG_"`
}

type UserConfig struct {
	AppName     string     `env:"APP_NAME" envDefault:"user-service"`
	GoEnv       string     `env:"GO_ENV" envDefault:"development" validate:"oneof=development production"`
	GRPCPort    string     `env:"USER_GRPC_PORT" envDefault:"50051"`
	MetricsPort string     `env:"USER_METRICS_PORT" envDefault:"9091"`
	Otel        OtelConfig `envPrefix:"OTEL_"`
	Log         LogConfig  `envPrefix:"LOG_"`
}

type CorsConfig struct {
	AllowOrigins string `env:"ALLOW_ORIGINS" envDefault:"*"`
}

type OtelConfig struct {
	ExporterEndpoint string `env:"EXPORTER_OTLP_ENDPOINT" envDefault:"localhost:4317"`
	ServiceName      string `env:"SERVICE_NAME" envDefault:"app"`
	TracesSampler    string `env:"TRACES_SAMPLER" envDefault:"always_on"`
}

type LogConfig struct {
	Level string `env:"LEVEL" envDefault:"info"`
}

func Load[T any]() T {
	var cfg T
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(cfg); err != nil {
		panic(err)
	}
	return cfg
}
