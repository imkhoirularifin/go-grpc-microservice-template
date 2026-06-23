package grpcutil

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	base := []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}
	return grpc.NewServer(append(base, opts...)...)
}

func Dial(ctx context.Context, host, port string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%s:%s", host, port)
	base := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	}
	conn, err := grpc.NewClient(target, append(base, opts...)...)
	if err != nil {
		return nil, fmt.Errorf("dial grpc %s: %w", target, err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return conn, nil
}
