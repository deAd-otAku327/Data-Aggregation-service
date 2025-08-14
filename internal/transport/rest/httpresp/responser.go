package httpresp

import (
	"data-aggregation-service/internal/types/dto"
	"encoding/json"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"
)

func MakeHTTPResponseJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func MakeHTTPErrorResponseJSON(w http.ResponseWriter, err *dto.ErrorResponse) {
	MakeHTTPResponseJSON(w, err.Code, err)
}
