package config

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppName   string
	HTTPAddr  string
	Workers   int
	QueueSize int
	LogLevel  slog.Level
}

func Load() Config {
	return Config{
		AppName:   getenv("APP_NAME", "go-worker"),
		HTTPAddr:  getenv("HTTP_ADDR", ":8080"),
		Workers:   getenvInt("WORKERS", 4),
		QueueSize: getenvInt("QUEUE_SIZE", 64),
		LogLevel:  parseLogLevel(getenv("LOG_LEVEL", "info")),
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func getenvInt(k string, def int) int {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
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
