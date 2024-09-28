package book

import (
	"casadocodigo/internal/category"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
		`,
		b.Title, b.Abstract, b.TableOfContent, b.Price, b.NumberOfPages, b.ISBN, b.PublishDate, b.Category.ID, b.AuthorID,
	).Scan(&b.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *bookRepository) FindByTitle(ctx context.Context, title string) (*Book, error) {
	var b Book
	var c category.Category

	var publishDate time.Time
	err := r.conn.QueryRow(
		ctx,
		`
		SELECT 
			b.id, 
			b.title,
			b.abstract,
			b.table_of_content,
			b.price,
			b.number_of_pages,
			b.isbn,
			b.publish_date,
			b.author_id,

			c.id,
			c.name,
			c.created_at
	 	FROM books b
		JOIN categories c ON b.category_id = c.id
		WHERE 
			title = $1
		`, title).Scan(
		&b.ID, &b.Title, &b.Abstract, &b.TableOfContent, &b.Price, &b.NumberOfPages, &b.ISBN, &publishDate, &b.AuthorID,
		&c.ID, &c.Name, &c.CreatedAt,
	)

	b.PublishDate = civil.DateOf(publishDate)
	b.Category = c

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}

	return &b, nil
}

func (r *bookRepository) List(ctx context.Context) ([]Book, error) {
	rows, err := r.conn.Query(
		ctx,
		`
		SELECT 
			b.id, 
			b.title,
			b.abstract,
			b.table_of_content,
			b.price,
			b.number_of_pages,
			b.isbn,
			b.publish_date,
			b.author_id,

			c.id,
			c.name,
			c.created_at
	 	FROM books b
		JOIN categories c ON b.category_id = c.id
		ORDER BY b.created_at DESC
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		var c category.Category

		var publishDate time.Time
		if err := rows.Scan(
			&b.ID, &b.Title, &b.Abstract, &b.TableOfContent, &b.Price, &b.NumberOfPages, &b.ISBN, &publishDate, &b.AuthorID,
			&c.ID, &c.Name, &c.CreatedAt,
		); err != nil {
			return nil, err
		}

		b.PublishDate = civil.DateOf(publishDate)
		b.Category = c

		books = append(books, b)
	}

	return books, nil
}
