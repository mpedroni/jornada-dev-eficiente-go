package category

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryAlreadyExists = errors.New("category with the given name already exists")
)

type CategoryService interface {
	Create(ctx context.Context, name string) (*Category, error)
	FindByName(ctx context.Context, name string) (*Category, error)
}

type categoryService struct {
	pool *pgxpool.Pool
}

func NewCategoryService(p *pgxpool.Pool) CategoryService {
	return &categoryService{
		pool: p,
	}
}

func (s *categoryService) Create(ctx context.Context, name string) (*Category, error) {
	existing, err := s.FindByName(ctx, name)
	if err != nil && !errors.Is(err, ErrCategoryNotFound) {
		return nil, err
	}

	if existing != nil {
		return nil, ErrCategoryAlreadyExists
	}

	c, err := NewCategory(name)
	if err != nil {
		return nil, err
	}

	err = s.pool.QueryRow(
		ctx,
		"INSERT INTO categories (name, created_at) VALUES ($1, $2) RETURNING ID",
		c.Name, c.CreatedAt,
	).Scan(&c.ID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *categoryService) FindByName(ctx context.Context, name string) (*Category, error) {
	c := &Category{}

	err := s.pool.QueryRow(
		ctx,
		"SELECT id, name, created_at FROM categories WHERE name = $1",
		name,
	).Scan(&c.ID, &c.Name, &c.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}

		return nil, err
	}

	return c, nil
}
