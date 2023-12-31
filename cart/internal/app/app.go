package app

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"route256/cart/internal/repository/postgres"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

type App struct {
	wg sync.WaitGroup
}

func (app *App) Run() error {
	// Init global context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Client connections
	lomsConn, productsConn, err := initClientConnections(ctx)
	if err != nil {
		return err
	}
	defer closeClientConnections(lomsConn, productsConn)

	// Clients
	lomsClient, productClient := initClients(lomsConn, productsConn)

	dbpool, err := getDBConnPool(ctx)
	if err != nil {
		return err
	}
	defer dbpool.Close()

	// Repository
	cartRepo := postgres.NewCartRepostiory(dbpool)

	// Services
	services := initServices(lomsClient, productClient, cartRepo)

	// Controller
	grpcServer, lis, err := initGRPCServer(services)
	if err != nil {
		return err
	}

	grpcGWServer, err := initGRPCGateway(ctx, lis)
	if err != nil {
		return err
	}

	go initGracefullShutdown(ctx, cancel, grpcServer, grpcGWServer)

	go app.StartGRPCServer(ctx, grpcServer, lis)

	app.StartGRPCGateway(grpcGWServer)

	app.Wait()

	return nil
}

func (app *App) StartGRPCServer(ctx context.Context, grpcServer *grpc.Server, lis net.Listener) {
	app.wg.Add(1)
	defer app.wg.Done()

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func (app *App) StartGRPCGateway(grcpGateway *http.Server) {
	log.Printf("Serving gRPC-Gateway on %s\n", grcpGateway.Addr)

	if err := grcpGateway.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("failed to serve: %v", err)
		}
	}
}

func (app *App) Wait() {
	app.wg.Wait()
}

func initGracefullShutdown(
	ctx context.Context,
	cancelGlobalCtx context.CancelFunc,
	grpcServer *grpc.Server,
	grpcGWServer *http.Server,
) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
	log.Println("Service gracefull shutdowning")

	cancelGlobalCtx()

	grpcServer.GracefulStop()

	if err := grpcGWServer.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
