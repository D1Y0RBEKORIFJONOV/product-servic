package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	"server/internal/gprc/protudct"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	gRPCServer *grpc.Server
	logger     *slog.Logger
	port       int
}

func New(logger *slog.Logger, port int, productServer protudct.Product) *App {
	grpcServer := grpc.NewServer()

	protudct.Register(grpcServer, productServer)
	reflection.Register(grpcServer)
	return &App{
		gRPCServer: grpcServer,
		logger:     logger,
		port:       port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.logger.With(
		slog.String("operation ", op),
		slog.Int("port", a.port),
	)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}
	defer l.Close()
	log.Info("grpc server is runing :", slog.String("adddres", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) ShutDown() {
	const op = "grpc.Shutdown"
	a.logger.With(slog.String("optarion", op)).Info("grpc is shutdinw down ", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
