package app

import (
	"log/slog"
	grpcapp "server/internal/app/grpc"
	"server/internal/infrastructure/repository/postgresql"
	configpkg "server/internal/pkg/config"
	"server/internal/pkg/postgres"
	"server/internal/services/products"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(logger *slog.Logger, grpcPort int, configS *configpkg.Config) *App {
	db, err := postgres.New(configS)
	if err != nil {
		panic(err)
	}
	storage := postgresql.NewProductRepository(db)
	product := products.NewProduct(logger, storage, storage, storage, storage)
	GRPCApp := grpcapp.New(logger, grpcPort, product)
	return &App{
		GRPCServer: GRPCApp,
	}
}
