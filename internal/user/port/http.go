package port

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moogu999/barito-be/internal/common/response"
	"github.com/moogu999/barito-be/internal/user/domain/entity"
	"github.com/moogu999/barito-be/internal/user/port/oapi"
	"github.com/moogu999/barito-be/internal/user/usecase"
)

func NewHandler(r chi.Router, svc usecase.User) http.Handler {
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
		res := oapi.ErrorResponse{
			Message: err.Error(),
		}

		if errors.Is(err, entity.ErrEmailIsUsed) {
			return oapi.CreateUser409JSONResponse(res), nil
		}

		return oapi.CreateUser500JSONResponse(res), nil
	}

	return oapi.CreateUser201Response(oapi.CreateUser201Response{}), nil
}

func (h *httpServer) CreateSession(ctx context.Context, request oapi.CreateSessionRequestObject) (oapi.CreateSessionResponseObject, error) {
	id, err := h.svc.CreateSession(ctx, string(request.Body.Email), request.Body.Password)
	if err != nil {
		res := oapi.ErrorResponse{
			Message: err.Error(),
		}

		if errors.Is(err, entity.ErrNotRegistered) {
			return oapi.CreateSession404JSONResponse(res), nil
		}
		if errors.Is(err, entity.ErrIncorrectPassword) {
			return oapi.CreateSession401JSONResponse(res), nil
		}

		return oapi.CreateSession500JSONResponse(res), nil
	}

	return oapi.CreateSession201JSONResponse(oapi.CreateSession201JSONResponse{
		Id: id,
	}), nil
}
