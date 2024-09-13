package infra

import (
	"casadocodigo/internal/author"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type authorRepository struct {
	conn *pgxpool.Pool
}

func NewAuthorRepository(pool *pgxpool.Pool) author.AuthorRepository {
	return &authorRepository{
		conn: pool,
	}
}

func (r *authorRepository) Save(ctx context.Context, a *author.Author) error {
	err := r.conn.QueryRow(
		ctx,
		"INSERT INTO authors (name, email, description, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
		a.Name,
		a.Email,
		a.Description,
		a.CreatedAt,
	).Scan(&a.ID)
	if err != nil {
		log.Printf("while saving author: %v\n", err)
		return err
	}

	return nil
}

func (r *authorRepository) FindByEmail(ctx context.Context, email string) (*author.Author, error) {
	a := author.Author{}

	err := r.conn.QueryRow(
		ctx,
		"SELECT id, name, email, description, created_at FROM authors WHERE email = $1",
		email,
	).Scan(&a.ID, &a.Name, &a.Email, &a.Description, &a.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, author.ErrAuthorNotFound
		}
		return nil, err
	}

	return &a, nil
}
