package port

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
	"github.com/moogu999/barito-be/internal/book/port/oapi"
	"github.com/moogu999/barito-be/internal/book/usecase"
	"github.com/moogu999/barito-be/internal/common/response"
)

func NewHandler(r chi.Router, svc usecase.BookUseCase) http.Handler {
	si := oapi.NewStrictHandlerWithOptions(&httpServer{svc}, nil, oapi.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  response.ErrorHandlerFunc(),
		ResponseErrorHandlerFunc: response.ErrorHandlerFunc(),
	})

	return oapi.HandlerFromMux(si, r)
}

type httpServer struct {
	svc usecase.BookUseCase
}

func (h *httpServer) FindBooks(ctx context.Context, request oapi.FindBooksRequestObject) (oapi.FindBooksResponseObject, error) {
	filter := repository.BookFilter{}

	if request.Params.Author != nil {
		filter.Author = *request.Params.Author
	}

	if request.Params.Title != nil {
		filter.Title = *request.Params.Title
	}

	data, err := h.svc.FindBooks(ctx, filter)
	if err != nil {
		return oapi.FindBooks500JSONResponse(oapi.ErrorResponse{
			Message: err.Error(),
		}), nil
	}

	books := make([]oapi.Book, 0)
	for _, val := range data {
		books = append(books, oapi.Book{
			Id:     val.ID,
			Title:  val.Title,
			Author: val.Author,
			Isbn:   val.ISBN,
			Price:  val.Price,
		})
	}

	return oapi.FindBooks200JSONResponse(oapi.FindBooks200JSONResponse{Books: books}), nil
}
