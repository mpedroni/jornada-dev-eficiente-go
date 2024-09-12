package main

import (
	"casadocodigo/internal/author"
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

	authorService := author.NewAuthorService(pool)
	authorHandler := author.NewAuthorHandler(authorService)

	r.POST("/authors", authorHandler.Create)

	r.Run()
}
