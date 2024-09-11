package author

import (
	"casadocodigo/internal/rest"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthorHandler interface {
	Create(c *gin.Context)
}

type authorHandler struct {
	authors AuthorService
}

func NewAuthorHandler(as AuthorService) AuthorHandler {
	return &authorHandler{
		authors: as,
	}
}

func (h *authorHandler) Create(c *gin.Context) {
	req := &CreateAuthorRequest{}

	rest.ValidateStruct(c, req)

	a, err := h.authors.Create(c.Request.Context(), req.Name, req.Email, req.Description)
	if err != nil {
		if err == ErrEmailAlreadyExists {
			c.JSON(400, map[string]string{"error": err.Error()})
			return
		}

		log.Printf("while creating author: %v\n", err)
		c.JSON(500, map[string]string{"error": "internal server error"})
		return
	}

	log.Printf("author with id %d created successfully\n", a.ID)
}
