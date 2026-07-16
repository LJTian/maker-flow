package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/LJTian/maker-flow/templates/apps/go-api/internal/config"
	"github.com/LJTian/maker-flow/templates/apps/go-api/internal/handler"
	"github.com/LJTian/maker-flow/templates/apps/go-api/internal/middleware"
)

type Server struct {
	cfg    config.Config
	logger *slog.Logger
	http   *http.Server
}

func New(cfg config.Config, logger *slog.Logger) *Server {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.Recover(logger))
	r.Use(middleware.RequestLogger(logger))
	r.Use(middleware.CORS())

	r.GET("/health", handler.Health)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handler.Ping)
	}

	return &Server{
		cfg:    cfg,
		logger: logger,
		http: &http.Server{
			Addr:         cfg.HTTPAddr,
			Handler:      r,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	s.logger.Info("starting server",
		"app", s.cfg.AppName,
		"env", s.cfg.AppEnv,
		"addr", s.cfg.HTTPAddr,
	)
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
