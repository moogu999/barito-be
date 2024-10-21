package port

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	bookEntity "github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/common/response"
	"github.com/moogu999/barito-be/internal/order/domain/entity"
	"github.com/moogu999/barito-be/internal/order/port/oapi"
	"github.com/moogu999/barito-be/internal/order/usecase"
	userEntity "github.com/moogu999/barito-be/internal/user/domain/entity"
	"github.com/oapi-codegen/runtime/types"
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
	cartItems := make([]usecase.CartItem, 0)
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

		if errors.Is(err, entity.ErrInvalidQuantity) {
			return oapi.CreateOrder422JSONResponse(res), nil
		}

		return oapi.CreateOrder500JSONResponse(res), nil
	}

	return oapi.CreateOrder201JSONResponse(oapi.CreateOrder201JSONResponse{
		Id: id,
	}), nil
}

func (h *httpServer) FindOrders(ctx context.Context, request oapi.FindOrdersRequestObject) (oapi.FindOrdersResponseObject, error) {
	orders, err := h.svc.FindOrders(ctx, request.Params.UserId)
	if err != nil {
		return oapi.FindOrders500JSONResponse(oapi.CreateOrder500JSONResponse{Message: err.Error()}), nil
	}

	res := make([]oapi.Order, 0)
	for _, val := range orders {
		items := make([]oapi.ItemResponse, 0)
		for _, item := range val.Items {
			items = append(items, oapi.ItemResponse{
				Id:     &item.ID,
				BookId: &item.BookID,
				Title:  &item.Title,
				Author: &item.Author,
				Qty:    &item.Qty,
				Price:  &item.Price,
			})
		}

		res = append(res, oapi.Order{
			Id:          val.ID,
			UserId:      val.UserID,
			Email:       types.Email(val.Email),
			Items:       items,
			TotalAmount: val.TotalAmount,
			CreatedAt:   val.CreatedAt,
		})
	}

	return oapi.FindOrders200JSONResponse(oapi.FindOrders200JSONResponse{Orders: res}), nil
}
