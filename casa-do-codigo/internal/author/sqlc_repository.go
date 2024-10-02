package author

import (
	"casadocodigo/casadocodigosqlc"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type sqlcAuthorRepository struct {
	queries *casadocodigosqlc.Queries
}

func NewSqlcAuthorRepository(q *casadocodigosqlc.Queries) AuthorRepository {
	return &sqlcAuthorRepository{
		queries: q,
	}
}

func (r *sqlcAuthorRepository) Save(ctx context.Context, a *Author) error {
	createdAt := pgtype.Timestamp{}
	createdAt.Scan(a.CreatedAt)
	id, err := r.queries.CreateAuthor(ctx, casadocodigosqlc.CreateAuthorParams{
		Name:        a.Name,
		Email:       a.Email,
		Description: &a.Description,
		CreatedAt:   createdAt,
	})
	if err != nil {
		return err
	}
	a.ID = int(id)
	return nil
}

func (r *sqlcAuthorRepository) FindByEmail(ctx context.Context, email string) (*Author, error) {
	a, err := r.queries.FindAuthorByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}
		return nil, err
	}

	return &Author{
		ID:          int(a.ID),
		Name:        a.Name,
		Email:       a.Email,
		Description: *a.Description,
		CreatedAt:   a.CreatedAt.Time,
	}, nil
}
