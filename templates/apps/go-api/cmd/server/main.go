package main

import (
	"log/slog"
	"os"

	"github.com/LJTian/maker-flow/templates/go-api/internal/config"
	"github.com/LJTian/maker-flow/templates/go-api/internal/server"
)

func main() {
	cfg := config.Load()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}))

	srv := server.New(cfg, logger)
	if err := srv.Run(); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
