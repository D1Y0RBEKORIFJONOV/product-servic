package main

import (
	"log/slog"
	"os"
	"os/signal"
	"server/internal/app"
	"server/internal/config"
	configpkg "server/internal/pkg/config"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting server",
		slog.Any("config", cfg))
	con, err := configpkg.NewConfig()
	if err != nil {
		log.Error(err.Error())
	}
	application := app.NewApp(log, cfg.GRPC.Port, con)

	go application.GRPCServer.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	log.Info("received shutdown signal", slog.String("signal", sig.String()))
	application.GRPCServer.ShutDown()
	log.Info("shutting down server")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
