package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/LJTian/maker-flow/templates/go-api/internal/config"
	"github.com/LJTian/maker-flow/templates/go-api/internal/handler"
	"github.com/LJTian/maker-flow/templates/go-api/internal/middleware"
)

type Server struct {
	cfg    config.Config
	logger *slog.Logger
	http   *http.Server
}

func New(cfg config.Config, logger *slog.Logger) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Recover(logger))
	r.Use(middleware.RequestLogger(logger))
	r.Use(middleware.CORS)

	r.Get("/health", handler.Health)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", handler.Ping)
	})

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
