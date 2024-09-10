package main

import (
	"casadocodigo/author"
	"casadocodigo/internal/rest"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateAuthorRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type HttpErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Detail    string    `json:"detail,omitempty"`
	Details   []string  `json:"details,omitempty"`
}

func CreateAuthorHandler(c *gin.Context, svc *author.Service) {
	req := &CreateAuthorRequest{}

	rest.ValidateStruct(c, req)

	svc.CreateAuthor(req.Name, req.Email, req.Description)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authors := author.NewService()

	r.POST("/authors", func(ctx *gin.Context) {
		CreateAuthorHandler(ctx, authors)
	})

	r.Run()
}
