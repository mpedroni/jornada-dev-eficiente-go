package category

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryRepository struct {
	conn *pgxpool.Pool
}

func NewPgxCategoryRepository(pool *pgxpool.Pool) CategoryRepository {
	return &categoryRepository{
		conn: pool,
	}
}

func (r *categoryRepository) Save(ctx context.Context, c *Category) error {
	err := r.conn.QueryRow(
		ctx,
		"INSERT INTO categories (name, created_at) VALUES ($1, $2) RETURNING id",
		c.Name, c.CreatedAt,
	).Scan(&c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) FindByName(ctx context.Context, name string) (*Category, error) {
	c := Category{}

	err := r.conn.QueryRow(
		context.Background(),
		"SELECT id, name, created_at FROM categories WHERE name = $1",
		name,
	).Scan(&c.ID, &c.Name, &c.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}

		return nil, err
	}

	return &c, nil
}
