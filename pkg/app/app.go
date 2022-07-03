package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/leaderseek/api-go/service"
	"github.com/leaderseek/service/pkg/server"
	"github.com/soheilhy/cmux"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	address     string
	logger      *zap.Logger
	grpc_server *grpc.Server
}

func NewApp(config *Config) (*App, error) {
	var app App
	app.address = config.HTTP.Address()

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config, error = %v", err)
	}

	logger, err := config.ZapLogger.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to start zap logger, error = %v", err)
	}

	app.logger = logger

	grpc_server := grpc.NewServer()
	reflection.Register(grpc_server)

	temporalOpts := client.Options{
		HostPort: config.TemporalClient.HostPort,
		Logger:   ZapToTemporalLogger(logger),
	}

	temporal, err := client.Dial(temporalOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to create temporal client, error = %v", err)
	}

	service.RegisterLeaderseekServer(grpc_server, server.NewServer(logger, temporal, config.ServerConfig))

	app.grpc_server = grpc_server

	return &app, nil
}

func (app *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	lisCfg := net.ListenConfig{}

	lis, err := lisCfg.Listen(ctx, "tcp", app.address)
	if err != nil {
		app.logger.Error("failed to create tcp listener", zap.Error(err))
		cancel()
		return
	}

	go func() {
		sigterm := <-termChan
		app.logger.Info("shutdown process initiated", zap.Any("sigterm", sigterm))
		cancel()
	}()

	mux := cmux.New(lis)
	grpcL := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	go func() {
		sErr := mux.Serve()
		if sErr != nil {
			app.logger.Fatal("failed to serve cmux", zap.Error(err))
		}
	}()

	app.logger.With(zap.Any("address", grpcL.Addr())).Info("serving grpc")

	go func() {
		sErr := app.grpc_server.Serve(grpcL)
		if sErr != nil {
			app.logger.Error("failed to serve grpc", zap.Error(err))
		}
	}()

	app.logger.Info("app started")
	<-ctx.Done()
	app.logger.Info("app shut down")
}
