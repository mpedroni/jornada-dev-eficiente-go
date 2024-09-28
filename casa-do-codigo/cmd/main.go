package main

import (
	"casadocodigo/internal/author"
	"casadocodigo/internal/book"
	"casadocodigo/internal/category"
	"casadocodigo/internal/database"
	"context"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var migrate bool
	var rollback bool

	flag.BoolVar(&migrate, "migrate", false, "run the migrations")
	flag.BoolVar(&rollback, "rollback", false, "run the rollback")
	flag.Parse()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pool, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/casa_do_codigo")
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	if rollback {
		database.Rollback(pool)
	}

	if migrate {
		database.Migrate(pool)
	}

	authorRepository := author.NewPgxAuthorRepository(pool)
	authorService := author.NewAuthorService(authorRepository)
	authorHandler := author.NewAuthorHandler(authorService)

	categoryRepository := category.NewPgxCategoryRepository(pool)
	categoryService := category.NewCategoryService(categoryRepository)
	categoryHandler := category.NewCategoryHandler(categoryService)

	bookRepository := book.NewPgxBookRepository(pool)
	bookService := book.NewBookService(bookRepository)
	bookHandler := book.NewBookHandler(bookService, categoryService)

	r.POST("/authors", authorHandler.Create)
	r.POST("/categories", categoryHandler.Create)
	r.POST("/books", bookHandler.Create)

	r.Run()
}
