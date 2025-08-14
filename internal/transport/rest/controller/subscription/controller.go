package subscription

import (
	"data-aggregation-service/internal/mappers/dtomap"
	"data-aggregation-service/internal/mappers/modelmap"
	"data-aggregation-service/internal/service/apierrors"
	"data-aggregation-service/internal/transport/rest/httperror"
	"data-aggregation-service/internal/transport/rest/httpresp"
	"data-aggregation-service/internal/types/domain"
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/validation"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

const URLParamSubID = "subId"

type subsControllerImpl struct {
	service    domain.SubscriptionService
	validation *validation.Validation
	logger     *slog.Logger
}

func New(s domain.SubscriptionService, v *validation.Validation, l *slog.Logger) domain.SubscriptionController {
	return &subsControllerImpl{
		service:    s,
		validation: v,
		logger:     l,
	}
}

func (c *subsControllerImpl) HandleCreateSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.CreateSubscriptionRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(
				[]string{apierrors.ErrParsingRequest.Error()}, http.StatusBadRequest))
			return
		}

		err = c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		response, err := c.service.CreateSubscription(r.Context(), modelmap.MapToSubscription(&request))
		if err != nil {
			code, apierr := httperror.ResolveHTTPErrorStatusCode(err, c.logger)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		httpresp.MakeHTTPResponseJSON(w, http.StatusCreated, response)
	}
}

func (c *subsControllerImpl) HandleGetSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.GetSubscriptionRequest{
			SubID: mux.Vars(r)[URLParamSubID],
		}

		err := c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		response, err := c.service.GetSubscription(r.Context(), modelmap.MapGetSubscriptionToSubscriptionID(&request))
		if err != nil {
			code, apierr := httperror.ResolveHTTPErrorStatusCode(err, c.logger)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		httpresp.MakeHTTPResponseJSON(w, http.StatusOK, response)
	}
}

func (c *subsControllerImpl) HandleUpdateSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.UpdateSubscriptionRequest{
			SubID: mux.Vars(r)[URLParamSubID],
		}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(
				[]string{apierrors.ErrParsingRequest.Error()}, http.StatusBadRequest))
			return
		}

		err = c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		err = c.service.UpdateSubscription(r.Context(),
			modelmap.MapUpdateSubscriptionToSubscriptionID(&request),
			modelmap.MapToSubscriptionPatch(&request),
		)
		if err != nil {
			code, apierr := httperror.ResolveHTTPErrorStatusCode(err, c.logger)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		httpresp.MakeHTTPResponseJSON(w, http.StatusNoContent, nil)
	}
}

func (c *subsControllerImpl) HandleDeleteSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.DeleteSubscriptionRequest{
			SubID: mux.Vars(r)[URLParamSubID],
		}
		err := c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		err = c.service.DeleteSubsription(r.Context(), modelmap.MapDeleteSubscriptionToSubscriptionID(&request))
		if err != nil {
			code, apierr := httperror.ResolveHTTPErrorStatusCode(err, c.logger)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		httpresp.MakeHTTPResponseJSON(w, http.StatusNoContent, nil)
	}
}

func (c *subsControllerImpl) HandleListSubscriptions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm() // Semicolon and other separators exclude "&"" generate errors.
		if err != nil {
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(
				[]string{apierrors.ErrParsingRequest.Error()}, http.StatusBadRequest))
			return
		}

		request := dto.ListSubscriptionsRequest{}
		err = schema.NewDecoder().Decode(&request, r.Form)
		if err != nil {
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierrors.ErrParsingRequest.Error()}, http.StatusBadRequest))
			return
		}

		err = c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		response, err := c.service.ListSubscriptions(r.Context(), modelmap.MapToSubscriptionFilterParams(&request))
		if err != nil {
			code, apierr := httperror.ResolveHTTPErrorStatusCode(err, c.logger)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		httpresp.MakeHTTPResponseJSON(w, http.StatusOK, response)
	}
}

func (c *subsControllerImpl) HandleGetSubscriptionsTotalCost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm() // Semicolon and other separators exclude "&"" generate errors.
		if err != nil {
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(
				[]string{apierrors.ErrParsingRequest.Error()}, http.StatusBadRequest))
			return
		}

		request := dto.GetTotalCostRequest{}
		err = schema.NewDecoder().Decode(&request, r.Form)
		if err != nil {
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(
				[]string{apierrors.ErrParsingRequest.Error()}, http.StatusBadRequest))
			return
		}

		err = c.validation.Validator.Struct(&request)
		if err != nil {
			msgs := validation.CollectValidationErrors(err, c.validation.Translator)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse(msgs, http.StatusBadRequest))
			return
		}

		response, err := c.service.GetSubscriptionsTotalCost(r.Context(), modelmap.MapToSubscriptionsTotalCostFilters(&request))
		if err != nil {
			code, apierr := httperror.ResolveHTTPErrorStatusCode(err, c.logger)
			httpresp.MakeHTTPErrorResponseJSON(w, dtomap.MapToErrorResponse([]string{apierr.Error()}, code))
			return
		}

		httpresp.MakeHTTPResponseJSON(w, http.StatusOK, response)
	}
}
