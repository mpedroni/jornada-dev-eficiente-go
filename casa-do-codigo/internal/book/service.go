package book

import (
	"casadocodigo/internal/category"
	"context"
	"errors"

	"cloud.google.com/go/civil"
)

var (
	ErrBookNotFound      = errors.New("book not found")
	ErrBookAlreadyExists = errors.New("book already exists")
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

	List(ctx context.Context) ([]Book, error)
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
	tableOfContent string,
	price float32,
	numberOfPages int,
	isbn string,
	publishDate civil.Date,
	category category.Category,
	authorID int,
) (*Book, error) {
	existing, err := s.books.FindByTitle(ctx, title)
	if err != nil && !errors.Is(err, ErrBookNotFound) {
		return nil, err
	}

	if existing != nil {
		return nil, ErrBookAlreadyExists
	}

	b, err := NewBook(
		title,
		abstract,
		tableOfContent,
		price,
		numberOfPages,
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

func (s *bookService) List(ctx context.Context) ([]Book, error) {
	books, err := s.books.List(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
