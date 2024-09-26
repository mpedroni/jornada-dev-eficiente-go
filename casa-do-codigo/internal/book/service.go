package book

import (
	"casadocodigo/internal/category"
	"context"

	"cloud.google.com/go/civil"
)

type BookService interface {
	Create(
		ctx context.Context,
		title,
		abstract,
		tableOfContents string,
		price float32,
		numberOfPages int,
		isbn string,
		publishDate civil.Date,
		category category.Category,
		authorID int,
	) (*Book, error)
}

type bookService struct {
	books BookRepository
}

func NewBookService(br BookRepository) BookService {
	return &bookService{
		books: br,
	}
}

func (s *bookService) Create(
	ctx context.Context,
	title,
	abstract,
	tableOfContents string,
	price float32,
	numberOfPages int,
	isbn string,
	publishDate civil.Date,
	category category.Category,
	authorID int,
) (*Book, error) {
	b, err := NewBook(
		title,
		abstract,
		tableOfContents,
		price,
		isbn,
		publishDate,
		category,
		authorID,
	)
	if err != nil {
		return nil, err
	}

	if err := s.books.Save(ctx, b); err != nil {
		return nil, err
	}

	return b, nil
}
