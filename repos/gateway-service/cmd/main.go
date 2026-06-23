package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/imkhoirularifin/gateway-service/internal/infrastructure"
	"github.com/imkhoirularifin/go-platform/pkg/config"
	"github.com/imkhoirularifin/go-platform/pkg/logger"
	"github.com/imkhoirularifin/go-platform/pkg/observability"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.Load[config.GatewayConfig]()
	logger.Setup(cfg.AppName, cfg.Log.Level, cfg.GoEnv)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	otelProvider, err := observability.Init(ctx, cfg.Otel.ServiceName, cfg.Otel.ExporterEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize observability")
	}
	otelProvider.StartMetricsServer(cfg.MetricsPort)

	server := infrastructure.NewServer(cfg)
	go server.Run()

	waitForShutdown(ctx, cancel, server, otelProvider)
}

func waitForShutdown(
	ctx context.Context,
	cancel context.CancelFunc,
	server *infrastructure.Server,
	otelProvider *observability.Provider,
) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Info().Msg("shutdown signal received")
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel()

	server.Shutdown(shutdownCtx)
	if err := otelProvider.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown observability")
	}
}
