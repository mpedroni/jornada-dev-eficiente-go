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
	books      BookService
	categories category.CategoryService
}

func NewBookHandler(bs BookService, cs category.CategoryService) BookHandler {
	return &bookHandler{
		books:      bs,
		categories: cs,
	}
}

func (h *bookHandler) Create(c *gin.Context) {
	req := &CreateBookRequest{}

	if !rest.ValidateStruct(c, req) {
		return
	}

	ctx := c.Request.Context()

	publishDate, _ := civil.ParseDate(req.PublishDate)
	category, err := h.categories.FindByName(ctx, req.Category)
	if err != nil {
		rest.BadRequest(c, fmt.Sprintf("category with name %s not found", req.Category))
		return
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
		*category,
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
