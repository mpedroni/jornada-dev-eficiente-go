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

	createBooksTableQuery := `CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		abstract TEXT NOT NULL,
		table_of_content TEXT NOT NULL,
		price FLOAT NOT NULL,
		number_of_pages INT NOT NULL,
		isbn VARCHAR(255) NOT NULL,
		publish_date DATE NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		category_id INT NOT NULL,
		author_id INT NOT NULL,

		FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE SET NULL,
		FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE
	)`

	migrations := []string{
		createAuthorsTableQuery,
		createCategoriesTableQuery,
		createBooksTableQuery,
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

	rollbackAuthorsTableQuery := `DROP TABLE IF EXISTS authors CASCADE`
	rollbackCategoriesTableQuery := "DROP TABLE IF EXISTS categories CASCADE"
	rollbackBooksTableQuery := "DROP TABLE IF EXISTS books"

	rollbacks := []string{
		rollbackCategoriesTableQuery,
		rollbackAuthorsTableQuery,
		rollbackBooksTableQuery,
	}

	for _, r := range rollbacks {
		if _, err := conn.Exec(ctx, r); err != nil {
			log.Fatalf("failed running rollback: %v", err)
		}
	}

	log.Println("rollback ran successfully")
}
