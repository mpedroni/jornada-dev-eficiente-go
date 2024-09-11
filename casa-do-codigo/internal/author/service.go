package author

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrAuthorNotFound     = errors.New("author not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type AuthorService interface {
	Create(ctx context.Context, name, email, description string) (*Author, error)
	FindByEmail(ctx context.Context, email string) (*Author, error)
}

type service struct {
	conn *pgxpool.Pool
}

func NewAuthorService(conn *pgxpool.Pool) AuthorService {
	return &service{
		conn: conn,
	}
}

func (s *service) Create(ctx context.Context, name, email, description string) (*Author, error) {
	existing, err := s.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, ErrAuthorNotFound) {
		log.Printf("while finding author by email: %v\n", err)
		return nil, err
	}

	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	a := NewAuthor(name, email, description)
	a.CreatedAt = time.Now()

	err = s.conn.QueryRow(
		ctx,
		"INSERT INTO authors (name, email, description, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
		a.Name,
		a.Email,
		a.Description,
		a.CreatedAt,
	).Scan(&a.ID)
	if err != nil {
		log.Printf("while saving author: %v\n", err)
		return nil, err
	}

	return &a, nil
}

func (s *service) FindByEmail(ctx context.Context, email string) (*Author, error) {
	a := Author{}

	err := s.conn.QueryRow(
		ctx,
		"SELECT id, name, email, description, created_at FROM authors WHERE email = $1",
		email,
	).Scan(&a.ID, &a.Name, &a.Email, &a.Description, &a.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}

		log.Printf("while scanning author: %v\n", err)
		return nil, err
	}

	return &a, nil
}
