package response

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(code int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func BuildResponseSuccess(code int, metaMessage string, metaStatus string, payload interface{}, w http.ResponseWriter) {
	response := Response{
		Meta: Meta{
			Message: metaMessage,
			Code:    code,
			Status:  metaStatus,
		},
		Data: payload,
	}

	respondWithJSON(code, response, w)
}

func BuildResponseFailure(code int, err interface{}, w http.ResponseWriter) {
	response := ResponseError{
		Meta: Meta{
			Message: "Error",
			Code:    code,
			Status:  "error",
		},
		Error: err,
	}

	respondWithJSON(code, response, w)

}
