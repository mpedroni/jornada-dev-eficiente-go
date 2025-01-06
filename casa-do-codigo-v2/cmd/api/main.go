package main

import (
	"cdc-v2/internal/author/database"
	"cdc-v2/internal/author/handler"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/cdc")
	if err != nil {
		panic(fmt.Errorf("opening database connection: %w", err))
	}

	if err := migrate(ctx, db); err != nil {
		panic(fmt.Errorf("migrating database: %w", err))
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("creating logger: %w", err))
	}

	authorRepository := database.New(db)
	authorHandler := handler.New(authorRepository, logger)

	r := gin.Default()

	authorHandler.RegisterRoutes(r)

	r.Run(":8080")
}

func migrate(ctx context.Context, db *pgx.Conn) error {
	dir, err := os.Getwd()

	schema, err := os.ReadFile(dir + "/db/schema.sql")
	if err != nil {
		return fmt.Errorf("opening db schema file: %w", err)
	}

	_, err = db.Exec(ctx, string(schema))
	if err != nil {
		return fmt.Errorf("executing db schema: %w", err)
	}
	return nil
}
