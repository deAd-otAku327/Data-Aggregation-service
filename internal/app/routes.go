package app

import (
	"data-aggregation-service/internal/transport/rest/controller"
	"data-aggregation-service/pkg/middleware"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func initRouting(controller controller.Controller, logger *slog.Logger) *mux.Router {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	router.Use(middleware.Logging(logger))

	router.HandleFunc("/subscriptions", controller.HandleCreateSubscription()).Methods(http.MethodPost)
	router.HandleFunc("/subscriptions/{subId}", controller.HandleCreateSubscription()).Methods(http.MethodGet)

	router.HandleFunc("/subscriptions/{subId}", controller.HandleUpdateSubscription()).Methods(http.MethodPatch)
	router.HandleFunc("/subscriptions/{subId}", controller.HandleDeleteSubscription()).Methods(http.MethodDelete)

	router.HandleFunc("/subscriptions", controller.HandleListSubscriptions()).Methods(http.MethodGet)
	router.HandleFunc("/total-cost", controller.HandleGetTotalCost()).Methods(http.MethodGet)

	return router
}
