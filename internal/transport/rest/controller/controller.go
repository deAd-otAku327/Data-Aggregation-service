package controller

import (
	"data-aggregation-service/internal/mappers/dtomap"
	"data-aggregation-service/internal/mappers/modelmap"
	"data-aggregation-service/internal/service"
	"data-aggregation-service/internal/transport/rest/responser"
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/validation"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Controller interface {
	HandleCreateSubscription() http.HandlerFunc
	HandleGetSubscription() http.HandlerFunc
	HandleUpdateSubscription() http.HandlerFunc
	HandleDeleteSubscription() http.HandlerFunc
	HandleListSubscriptions() http.HandlerFunc
	HandleGetSubscriptionsTotalCost() http.HandlerFunc
}

const URLParamSubID = "subId"

type controller struct {
	service    service.Service
	validation *validation.Validation
	logger     *slog.Logger
}

func New(s service.Service, v *validation.Validation, l *slog.Logger) Controller {
	return &controller{
		service:    s,
		validation: v,
		logger:     l,
	}
}

func (c *controller) HandleCreateSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.CreateSubscriptionRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{ErrParsingRequest.Error()}, http.StatusBadRequest))
			return
		}

		err = c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		response, err := c.service.CreateSubscription(r.Context(), modelmap.MapToSubscription(&request))
		if err != nil {
			code, apierr := resolveError(err, c.logger)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		responser.MakeResponseJSON(w, http.StatusCreated, response)
	}
}

func (c *controller) HandleGetSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.GetSubscriptionRequest{
			SubID: mux.Vars(r)[URLParamSubID],
		}
		err := c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		response, err := c.service.GetSubscription(r.Context(), modelmap.MapGetSubscriptionToSubscriptionID(&request))
		if err != nil {
			code, apierr := resolveError(err, c.logger)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, response)
	}
}

func (c *controller) HandleUpdateSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.UpdateSubscriptionRequest{
			SubID: mux.Vars(r)[URLParamSubID],
		}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{ErrParsingRequest.Error()}, http.StatusBadRequest))
			return
		}

		err = c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		err = c.service.UpdateSubscription(r.Context(),
			modelmap.MapUpdateSubscriptionToSubscriptionID(&request),
			modelmap.MapToSubscriptionPatch(&request),
		)
		if err != nil {
			code, apierr := resolveError(err, c.logger)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		responser.MakeResponseJSON(w, http.StatusNoContent, nil)
	}
}

func (c *controller) HandleDeleteSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.DeleteSubscriptionRequest{
			SubID: mux.Vars(r)[URLParamSubID],
		}
		err := c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		err = c.service.DeleteSubsription(r.Context(), modelmap.MapDeleteSubscriptionToSubscriptionID(&request))
		if err != nil {
			code, apierr := resolveError(err, c.logger)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		responser.MakeResponseJSON(w, http.StatusNoContent, nil)
	}
}

func (c *controller) HandleListSubscriptions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {

			return
		}

		request := dto.ListSubscriptionsRequest{}
		err = schema.NewDecoder().Decode(&request, r.Form)
		if err != nil {

			return
		}

		err = c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		response, err := c.service.ListSubscriptions(r.Context(), modelmap.MapToSubscriptionFilterParams(&request))
		if err != nil {

			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, response)
	}
}

func (c *controller) HandleGetSubscriptionsTotalCost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {

			return
		}

		request := dto.GetTotalCostRequest{}
		err = schema.NewDecoder().Decode(&request, r.Form)
		if err != nil {

			return
		}

		err = c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			responser.MakeErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		response, err := c.service.GetSubscriptionsTotalCost(r.Context(), modelmap.MapToSubscriptionsTotalCostFilters(&request))
		if err != nil {

			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, response)
	}
}
