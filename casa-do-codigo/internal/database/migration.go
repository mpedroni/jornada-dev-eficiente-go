package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(conn *pgxpool.Pool) {
	ctx := context.Background()

	createAuthorsTableQuery := `CREATE TABLE IF NOT EXISTS authors (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP NOT NULL
	)
	`

	if _, err := conn.Exec(ctx, createAuthorsTableQuery); err != nil {
		log.Fatalf("failed running migrations: %v", err)
	}

	log.Println("migrations ran successfully")
}

func Rollback(conn *pgxpool.Pool) {
	ctx := context.Background()

	rollbackAuthorsTableQuery := `DROP TABLE IF EXISTS authors`

	if _, err := conn.Exec(ctx, rollbackAuthorsTableQuery); err != nil {
		log.Fatalf("failed running rollback: %v", err)
	}

	log.Println("rollback ran successfully")
}
