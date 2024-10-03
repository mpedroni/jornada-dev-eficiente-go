package category

import (
	"casadocodigo/casadocodigosqlc"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type sqlcCategoryRepository struct {
	queries *casadocodigosqlc.Queries
}

func NewSqlcCategoryRepository(q *casadocodigosqlc.Queries) CategoryRepository {
	return &sqlcCategoryRepository{
		queries: q,
	}
}

func (r *sqlcCategoryRepository) Save(ctx context.Context, c *Category) error {
	id, err := r.queries.CreateCategory(ctx, casadocodigosqlc.CreateCategoryParams{
		Name:      c.Name,
		CreatedAt: pgtype.Timestamp{Time: c.CreatedAt, Valid: true},
	})
	if err != nil {
		return err
	}
	c.ID = int(id)

	return nil
}

func (r *sqlcCategoryRepository) FindByName(ctx context.Context, name string) (*Category, error) {
	category, err := r.queries.FindCategoryByName(ctx, name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	return &Category{
		ID:        int(category.ID),
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Time,
	}, nil
}
