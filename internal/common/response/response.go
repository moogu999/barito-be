package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

// ErrorHandlerFunc provides a default http handler for uncaptured
// HTTP errors.
func ErrorHandlerFunc() func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		if err == nil {
			err = errors.New("undefined behaviour")
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
	}
}
