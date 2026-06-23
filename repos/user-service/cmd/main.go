package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/imkhoirularifin/go-platform/pkg/config"
	"github.com/imkhoirularifin/go-platform/pkg/grpcutil"
	"github.com/imkhoirularifin/go-platform/pkg/logger"
	"github.com/imkhoirularifin/go-platform/pkg/observability"
	"github.com/imkhoirularifin/user-service/internal/handler"
	"github.com/imkhoirularifin/user-service/internal/service"
	userv1 "github.com/imkhoirularifin/proto-contracts/gen/go/user/v1"
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
