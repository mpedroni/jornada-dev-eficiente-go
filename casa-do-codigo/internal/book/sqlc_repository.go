package book

import (
	"casadocodigo/casadocodigosqlc"
	"casadocodigo/internal/category"
	"context"
	"errors"
	"time"

	"cloud.google.com/go/civil"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type sqlcBookRepository struct {
	queries *casadocodigosqlc.Queries
}

func NewSqlcBookRepository(q *casadocodigosqlc.Queries) BookRepository {
	return &sqlcBookRepository{
		queries: q,
	}
}

func (r *sqlcBookRepository) Save(ctx context.Context, book *Book) error {
	categoryID := int32(book.Category.ID)
	id, err := r.queries.CreateBook(ctx, casadocodigosqlc.CreateBookParams{
		Title:          book.Title,
		Abstract:       book.Abstract,
		TableOfContent: book.TableOfContent,
		Price:          float64(book.Price),
		NumberOfPages:  int32(book.NumberOfPages),
		Isbn:           book.ISBN,
		PublishDate:    pgtype.Date{Time: book.PublishDate.In(time.UTC), Valid: true},
		CategoryID:     &categoryID,
		AuthorID:       int32(book.AuthorID),
	})
	if err != nil {
		return err
	}

	book.ID = int(id)
	return nil
}

func (r *sqlcBookRepository) FindByTitle(ctx context.Context, title string) (*Book, error) {
	book, err := r.queries.FindBookByTitle(ctx, title)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}

	return &Book{
		ID:             int(book.ID),
		Title:          book.Title,
		Abstract:       book.Abstract,
		TableOfContent: book.TableOfContent,
		Price:          float32(book.Price),
		NumberOfPages:  int(book.NumberOfPages),
		ISBN:           book.Isbn,
		PublishDate:    civil.DateOf(book.PublishDate.Time),
		CreatedAt:      book.CreatedAt.Time,
		AuthorID:       int(book.AuthorID),
		Category: category.Category{
			ID:        int(*book.CategoryID),
			Name:      book.Category.Name,
			CreatedAt: book.Category.CreatedAt.Time,
		},
	}, nil
}

func (r *sqlcBookRepository) List(ctx context.Context) ([]Book, error) {
	result, err := r.queries.ListBooks(ctx)
	if err != nil {
		return nil, err
	}

	books := make([]Book, len(result))
	for i, b := range result {
		books[i] = Book{
			ID:             int(b.ID),
			Title:          b.Title,
			Abstract:       b.Abstract,
			TableOfContent: b.TableOfContent,
			Price:          float32(b.Price),
			NumberOfPages:  int(b.NumberOfPages),
			ISBN:           b.Isbn,
			PublishDate:    civil.DateOf(b.PublishDate.Time),
			CreatedAt:      b.CreatedAt.Time,
			AuthorID:       int(b.AuthorID),
			Category: category.Category{
				ID:        int(*b.CategoryID),
				Name:      b.Category.Name,
				CreatedAt: b.Category.CreatedAt.Time,
			},
		}
	}

	return books, nil
}
