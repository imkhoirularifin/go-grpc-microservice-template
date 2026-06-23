package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	userv1 "github.com/imkhoirularifin/go-grpc-microservice-template/gen/go/user/v1"
	"github.com/imkhoirularifin/go-grpc-microservice-template/internal/user/handler"
	"github.com/imkhoirularifin/go-grpc-microservice-template/internal/user/service"
	"github.com/imkhoirularifin/go-grpc-microservice-template/pkg/config"
	"github.com/imkhoirularifin/go-grpc-microservice-template/pkg/grpcutil"
	"github.com/imkhoirularifin/go-grpc-microservice-template/pkg/logger"
	"github.com/imkhoirularifin/go-grpc-microservice-template/pkg/observability"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.Load[config.UserConfig]()
	logger.Setup(cfg.AppName, cfg.Log.Level, cfg.GoEnv)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	otelProvider, err := observability.Init(ctx, cfg.Otel.ServiceName, cfg.Otel.ExporterEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize observability")
	}
	otelProvider.StartMetricsServer(cfg.MetricsPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	userService := service.NewUserService()
	grpcServer := grpcutil.NewServer()
	userv1.RegisterUserServiceServer(grpcServer, handler.NewUserHandler(userService))

	go func() {
		log.Info().Str("port", cfg.GRPCPort).Msg("user gRPC server started")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("gRPC server failed")
		}
	}()

	waitForShutdown(ctx, cancel, grpcServer, otelProvider)
}

func waitForShutdown(
	ctx context.Context,
	cancel context.CancelFunc,
	grpcServer interface{ GracefulStop() },
	otelProvider *observability.Provider,
) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Info().Msg("shutdown signal received")
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel()

	grpcServer.GracefulStop()
	if err := otelProvider.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown observability")
	}
}
