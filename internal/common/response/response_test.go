package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestErrorHandlerFunc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		fn           func(w http.ResponseWriter, r *http.Request, err error)
		err          error
		wantResponse ErrorResponse
	}{
		{
			name: "err is not nil",
			fn:   ErrorHandlerFunc(),
			err:  errors.New("unexpected error"),
			wantResponse: ErrorResponse{
				Message: "unexpected error",
			},
		},
		{
			name: "err is nil",
			fn:   ErrorHandlerFunc(),
			err:  nil,
			wantResponse: ErrorResponse{
				Message: "undefined behaviour",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/", nil)

			rr := httptest.NewRecorder()

			tt.fn(rr, req, tt.err)

			var res ErrorResponse
			json.NewDecoder(rr.Body).Decode(&res)

			if !strings.EqualFold(res.Message, tt.wantResponse.Message) {
				t.Errorf("different error message, got=%v, want=%v", res.Message, tt.wantResponse.Message)
			}
		})
	}
}
