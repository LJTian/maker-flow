package config

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	AppName  string
	AppEnv   string
	HTTPAddr string
	LogLevel slog.Level
}

func Load() Config {
	level := parseLogLevel(getenv("LOG_LEVEL", "info"))
	return Config{
		AppName:  getenv("APP_NAME", "go-api"),
		AppEnv:   getenv("APP_ENV", "development"),
		HTTPAddr: getenv("HTTP_ADDR", ":8080"),
		LogLevel: level,
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseLogLevel(raw string) slog.Level {
	switch strings.ToLower(raw) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
