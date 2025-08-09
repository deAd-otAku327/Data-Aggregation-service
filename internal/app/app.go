package app

import (
	"context"
	"data-aggregation-service/internal/config"
	"data-aggregation-service/internal/repository"
	"data-aggregation-service/internal/service"
	"data-aggregation-service/internal/transport/rest/controller"
	"fmt"
	"log/slog"
	"net/http"
)

type App struct {
	Server *http.Server
}

func New(cfg *config.Config) (*App, error) {
	repo := repository.New(cfg.PostgresDB)
	service := service.New(repo.Postgres)
	controller := controller.New(service)

	return &App{
		Server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
			Handler: initRouting(controller),
		},
	}, nil
}

func (s *App) Run() error {
	slog.Info("app starting on", slog.String("address", s.Server.Addr))
	return s.Server.ListenAndServe()
}

func (s *App) Shutdown() error {
	slog.Info("app shutting down...")
	return s.Server.Shutdown(context.Background())
}
