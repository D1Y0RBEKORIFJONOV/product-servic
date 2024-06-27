package app

import (
	"log/slog"
	grpcapp "server/internal/app/grpc"
)


type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(loggger *slog.Logger, grpcPort int,storagePath string) *App {
	panic("weg")
}