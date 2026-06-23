package infrastructure

import (
	"context"
	"fmt"

	userv1 "github.com/imkhoirularifin/go-grpc-microservice-template/gen/go/user/v1"
	"github.com/imkhoirularifin/go-grpc-microservice-template/internal/gateway/handler"
	"github.com/imkhoirularifin/go-grpc-microservice-template/lib/common"
	"github.com/imkhoirularifin/go-grpc-microservice-template/pkg/config"
	"github.com/imkhoirularifin/go-grpc-microservice-template/pkg/grpcutil"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

type Server struct {
	app *fiber.App
	cfg config.GatewayConfig
}

func NewServer(cfg config.GatewayConfig) *Server {
	app := fiber.New(fiber.Config{
		AppName:      cfg.AppName,
		ErrorHandler: common.ErrorHandler,
	})

	app.Use(otelfiber.Middleware())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.Cors.AllowOrigins,
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger:          &log.Logger,
		Fields:          []string{"latency", "status", "method", "url"},
		FieldsSnakeCase: true,
		SkipURIs:        []string{"/healthz", "/readyz"},
	}))

	conn, err := grpcutil.Dial(context.Background(), cfg.UserGRPCHost, cfg.UserGRPCPort)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to user gRPC service")
	}

	userClient := userv1.NewUserServiceClient(conn)
	userHandler := handler.NewUserHandler(userClient)

	api := app.Group("/api/v1")
	handler.NewHealthHandler(api)
	userHandler.Register(api.Group("/users"))

	app.Use(common.NotFoundHandler)

	return &Server{app: app, cfg: cfg}
}

func (s *Server) Run() {
	log.Info().Str("port", s.cfg.Port).Msg("gateway started")
	if err := s.app.Listen(fmt.Sprintf(":%s", s.cfg.Port)); err != nil {
		log.Fatal().Err(err).Msg("gateway failed")
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
