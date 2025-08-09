package controller

import (
	"data-aggregation-service/internal/service"
	"log/slog"
	"net/http"
)

type Controller interface {
	HandleCreateSubscription() http.HandlerFunc
	HandleGetSubscription() http.HandlerFunc
	HandleUpdateSubscription() http.HandlerFunc
	HandleDeleteSubscription() http.HandlerFunc
	HandleListSubscriptions() http.HandlerFunc
	HandleGetTotalCost() http.HandlerFunc
}

type controller struct {
	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, l *slog.Logger) Controller {
	return &controller{
		service: s,
		logger:  l,
	}
}

func (c *controller) HandleCreateSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c *controller) HandleGetSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c *controller) HandleUpdateSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c *controller) HandleDeleteSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c *controller) HandleListSubscriptions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c *controller) HandleGetTotalCost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
