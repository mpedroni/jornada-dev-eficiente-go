package book

import (
	"casadocodigo/internal/category"
	"context"
	"errors"
	"time"

	"cloud.google.com/go/civil"
)

type BookRepository interface {
	Save(ctx context.Context, book *Book) error
}

type ValidationError struct {
	errors []error
}

func (ve ValidationError) Error() string {
	result := ""
	for _, err := range ve.errors {
		result += err.Error() + "\n"
	}

	return result
}

func (ve *ValidationError) Add(err error) {
	ve.errors = append(ve.errors, err)
}

func (ve *ValidationError) AddString(err string) {
	ve.errors = append(ve.errors, errors.New(err))
}

type Book struct {
	ID             int
	Title          string
	Abstract       string
	TableOfContent string
	Price          float32
	NumberOfPages  int
	ISBN           string
	PublishDate    civil.Date
	Category       category.Category
	AuthorID       int
	CreatedAt      time.Time
}

func NewBook(
	title,
	abstract,
	tableOfContents string,
	price float32,
	isbn string,
	publishDate civil.Date,
	category category.Category,
	authorID int,

) (*Book, error) {
	b := &Book{
		ID:             -1,
		Title:          title,
		Abstract:       abstract,
		TableOfContent: tableOfContents,
		Price:          price,
		ISBN:           isbn,
		PublishDate:    publishDate,
		Category:       category,
		AuthorID:       authorID,
		CreatedAt:      time.Now(),
	}

	if err := b.Validate(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Book) Validate() error {
	var ve ValidationError

	if b.Title == "" {
		ve.AddString("title is required")
	}

	if b.Abstract == "" {
		ve.AddString("abstract is required")
	}

	if b.TableOfContent == "" {
		ve.AddString("table of content is required")
	}

	if b.Price < 20 {
		ve.AddString("price must be at least 20")
	}

	if b.NumberOfPages < 100 {
		ve.AddString("number of pages must be at least 100")
	}

	if b.ISBN == "" {
		ve.AddString("isbn is required")
	}

	if b.PublishDate.IsZero() {
		ve.AddString("publish date is required")
	}

	if b.PublishDate.Before(civil.DateOf(time.Now())) {
		ve.AddString("publish date must be in the future")
	}

	if b.Category.ID == 0 {
		ve.AddString("category is required")
	}

	if b.AuthorID == 0 {
		ve.AddString("author is required")
	}

	if len(ve.errors) > 0 {
		return ve
	}

	return nil
}
