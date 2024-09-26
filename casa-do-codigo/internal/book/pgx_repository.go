package book

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type bookRepository struct {
	conn *pgxpool.Pool
}

func NewPgxBookRepository(conn *pgxpool.Pool) BookRepository {
	return &bookRepository{
		conn: conn,
	}
}

func (r *bookRepository) Save(ctx context.Context, b *Book) error {
	err := r.conn.QueryRow(
		ctx,
		`INSERT INTO books 
			(
				title,
				abstract,
				table_of_content,
				price,
				number_of_pages,
				isbn,
				publish_date,
				category_id,
				author_id
			)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		b.Title, b.Abstract, b.TableOfContent, b.Price, b.NumberOfPages, b.ISBN, b.PublishDate, b.Category.ID, b.AuthorID,
	).Scan(&b.ID)
	if err != nil {
		return err
	}

	return nil
}
