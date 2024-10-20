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
	data, err := h.svc.FindBooks(ctx, repository.BookFilter{
		Author: *request.Params.Author,
		Title:  *request.Params.Title,
	})
	if err != nil {
		return oapi.FindBooks500JSONResponse(oapi.ErrorResponse{
			Message: err.Error(),
		}), nil
	}

	books := make([]oapi.Book, len(data))
	for _, val := range data {
		books = append(books, oapi.Book{
			Id:     val.ID,
			Title:  val.Title,
			Author: val.Author,
			Isbn:   val.ISBN,
		})
	}

	return oapi.FindBooks200JSONResponse(oapi.FindBooks200JSONResponse{Books: books}), nil
}
