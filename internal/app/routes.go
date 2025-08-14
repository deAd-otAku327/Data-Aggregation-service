package app

import (
	"data-aggregation-service/internal/types/domain"
	"data-aggregation-service/pkg/middleware"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func initRouting(controller domain.SubscriptionController, logger *slog.Logger) *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	router.Use(middleware.Logging(logger))

	router.HandleFunc("/subscriptions", controller.HandleCreateSubscription()).Methods(http.MethodPost)
	router.HandleFunc("/subscriptions/{subId}", controller.HandleGetSubscription()).Methods(http.MethodGet)

	router.HandleFunc("/subscriptions/{subId}", controller.HandleUpdateSubscription()).Methods(http.MethodPatch)
	router.HandleFunc("/subscriptions/{subId}", controller.HandleDeleteSubscription()).Methods(http.MethodDelete)

	router.HandleFunc("/subscriptions", controller.HandleListSubscriptions()).Methods(http.MethodGet)
	router.HandleFunc("/subscriptions/cost/total", controller.HandleGetSubscriptionsTotalCost()).Methods(http.MethodGet)

	return router
}
