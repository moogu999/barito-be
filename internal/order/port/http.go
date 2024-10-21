package port

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	bookEntity "github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/common/response"
	"github.com/moogu999/barito-be/internal/order/port/oapi"
	"github.com/moogu999/barito-be/internal/order/usecase"
	userEntity "github.com/moogu999/barito-be/internal/user/domain/entity"
)

func NewHandler(r chi.Router, svc usecase.OrderUseCase) http.Handler {
	si := oapi.NewStrictHandlerWithOptions(&httpServer{svc}, nil, oapi.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  response.ErrorHandlerFunc(),
		ResponseErrorHandlerFunc: response.ErrorHandlerFunc(),
	})

	return oapi.HandlerFromMux(si, r)
}

type httpServer struct {
	svc usecase.OrderUseCase
}

func (h *httpServer) CreateOrder(ctx context.Context, request oapi.CreateOrderRequestObject) (oapi.CreateOrderResponseObject, error) {
	cartItems := make([]usecase.CartItem, len(request.Body.Items))
	for _, val := range request.Body.Items {
		cartItems = append(cartItems, usecase.CartItem{
			BookID: val.BookId,
			Qty:    val.Qty,
		})
	}

	id, err := h.svc.CreateOrder(ctx, request.Body.UserId, cartItems)
	if err != nil {
		res := oapi.ErrorResponse{
			Message: err.Error(),
		}

		if errors.Is(err, userEntity.ErrUserNotFound) || errors.Is(err, bookEntity.ErrBooksNotFound) {
			return oapi.CreateOrder404JSONResponse(res), nil
		}

		return oapi.CreateOrder500JSONResponse(res), nil
	}

	return oapi.CreateOrder201JSONResponse(oapi.CreateOrder201JSONResponse{
		Id: id,
	}), nil
}
