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
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		description TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`

	createCategoriesTableQuery := `CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`

	migrations := []string{
		createAuthorsTableQuery,
		createCategoriesTableQuery,
	}

	for _, m := range migrations {
		if _, err := conn.Exec(ctx, m); err != nil {
			log.Fatalf("failed running migrations: %v", err)
		}
	}

	log.Println("migrations ran successfully")
}

func Rollback(conn *pgxpool.Pool) {
	ctx := context.Background()

	rollbackAuthorsTableQuery := `DROP TABLE IF EXISTS authors`
	rollbackCategoriesTableQuery := "DROP TABLE IF EXISTS categories"

	rollbacks := []string{
		rollbackCategoriesTableQuery,
		rollbackAuthorsTableQuery,
	}

	for _, r := range rollbacks {
		if _, err := conn.Exec(ctx, r); err != nil {
			log.Fatalf("failed running rollback: %v", err)
		}
	}

	log.Println("rollback ran successfully")
}
