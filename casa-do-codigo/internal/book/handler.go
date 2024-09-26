package book

import (
	"casadocodigo/internal/category"
	"casadocodigo/internal/rest"
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
	category := category.Category{}

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
		fmt.Println(err)
		return
	}
}
