package book

import (
	"casadocodigo/internal/category"
	"casadocodigo/internal/rest"
	"errors"
	"fmt"

	"cloud.google.com/go/civil"
	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	Create(c *gin.Context)
}

type bookHandler struct {
	books BookService
}

func NewBookHandler(bs BookService) BookHandler {
	return &bookHandler{
		books: bs,
	}
}

func (h *bookHandler) Create(c *gin.Context) {
	req := &CreateBookRequest{}

	if !rest.ValidateStruct(c, req) {
		return
	}

	publishDate, _ := civil.ParseDate(req.PublishDate)
	category := category.Category{
		ID: 1,
	}

	if _, err := h.books.Create(
		c.Request.Context(),
		req.Title,
		req.Abstract,
		req.TableOfContent,
		req.Price,
		req.NumberOfPages,
		req.ISBN,
		publishDate,
		category,
		req.AuthorID,
	); err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			rest.BadRequest(c, err.Error())
			return
		}

		if err == ErrBookAlreadyExists {
			rest.BadRequest(c, fmt.Sprintf("book with title %s already exists", req.Title))
			return
		}

		rest.InternalServerError(c)
		return
	}
}
