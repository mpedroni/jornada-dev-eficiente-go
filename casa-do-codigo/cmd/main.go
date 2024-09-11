package main

import (
	"casadocodigo/internal/author"
	"casadocodigo/internal/rest"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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

func CreateAuthorHandler(c *gin.Context, svc author.AuthorService) {
	req := &CreateAuthorRequest{}

	rest.ValidateStruct(c, req)

	svc.Create(c.Request.Context(), req.Name, req.Email, req.Description)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pool, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/casa_do_codigo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	if err := pool.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "unable to ping database: %v\n", err)
		os.Exit(1)
	}

	authorService := author.NewAuthorService(pool)
	authorHandler := author.NewAuthorHandler(authorService)

	r.POST("/authors", authorHandler.Create)

	r.Run()
}
