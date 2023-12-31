package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"route256/cart/pkg/middleware/logging"
	"route256/cart/pkg/middleware/panic"
	api "route256/loms/internal/api/loms"
	desc "route256/loms/pkg/api/loms/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func initGRPCServer(services *api.Deps) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			panic.Interceptor,
			logging.Interceptor,
		),
	)

	reflection.Register(grpcServer)

	controller := api.NewServer(services)

	desc.RegisterLOMSServer(grpcServer, controller)

	return grpcServer, lis, nil
}

func initGRPCGateway(ctx context.Context, lis net.Listener) (*http.Server, error) {
	mux := runtime.NewServeMux()

	if err := desc.RegisterLOMSHandlerFromEndpoint(
		ctx,
		mux,
		lis.Addr().String(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return nil, fmt.Errorf("register carts handler: %v", err)
	}

	return &http.Server{
		Addr:    os.Getenv("GRPC_GW_ADDR"),
		Handler: logging.WithHTTPLoggingMiddleware(mux),
	}, nil
}
