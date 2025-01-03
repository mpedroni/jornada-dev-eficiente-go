package database

import (
	"cdc-v2/internal/author/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) domain.AuthorRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, author *domain.Author) error {
	err := r.db.QueryRow(
		ctx,
		"INSERT INTO authors (name, email, description) VALUES ($1, $2, $3) RETURNING id",
		author.Name, author.Email, author.Description,
	).Scan(&author.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*domain.Author, error) {
	author := &domain.Author{}
	err := r.db.QueryRow(
		ctx,
		"SELECT id, name, email, description, created_at FROM authors WHERE email = $1",
		email,
	).Scan(&author.ID, &author.Name, &author.Email, &author.Description, &author.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrAuthorNotFound
		}
		return nil, err
	}

	return author, nil
}
