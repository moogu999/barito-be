package user

import (
	"context"
	"net/http"

	"github.com/moogu999/barito-be/internal/common/response"
	"github.com/moogu999/barito-be/internal/infra/port/user/oapi"
	"github.com/moogu999/barito-be/internal/usecase/user"
)

func NewHTTP(svc user.Service) *http.ServeMux {
	si := oapi.NewStrictHandlerWithOptions(&httpServer{svc}, nil, oapi.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  response.ErrorHandlerFunc(),
		ResponseErrorHandlerFunc: response.ErrorHandlerFunc(),
	})

	mux := http.NewServeMux()
	oapi.HandlerFromMux(si, mux)
	return mux
}

type httpServer struct {
	svc user.Service
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
