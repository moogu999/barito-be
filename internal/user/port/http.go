package port

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moogu999/barito-be/internal/common/response"
	"github.com/moogu999/barito-be/internal/user/port/oapi"
	"github.com/moogu999/barito-be/internal/user/usecase"
)

func NewHTTP(r chi.Router, svc usecase.User) http.Handler {
	si := oapi.NewStrictHandlerWithOptions(&httpServer{svc}, nil, oapi.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  response.ErrorHandlerFunc(),
		ResponseErrorHandlerFunc: response.ErrorHandlerFunc(),
	})

	return oapi.HandlerFromMux(si, r)
}

type httpServer struct {
	svc usecase.User
}

func (h *httpServer) CreateUser(ctx context.Context, request oapi.CreateUserRequestObject) (oapi.CreateUserResponseObject, error) {
	err := h.svc.CreateUser(ctx, string(request.Body.Email), request.Body.Password)
	if err != nil {
		return oapi.CreateUser500JSONResponse(oapi.ErrorResponse{
			Message: err.Error(),
		}), nil
	}

	return oapi.CreateUser201Response(oapi.CreateUser201Response{}), nil
}
