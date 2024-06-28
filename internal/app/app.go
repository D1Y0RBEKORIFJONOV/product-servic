package app

import (
	"log/slog"
	grpcapp "server/internal/app/grpc"
	configpkg "server/internal/pkg/config"
	"server/internal/pkg/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(loggger *slog.Logger, grpcPort int, configS *configpkg.Config) *App {
	postgres.New(configS)
	panic("f")
}
