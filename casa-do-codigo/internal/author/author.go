package author

import (
	"context"
	"fmt"
	"time"
)

type AuthorRepository interface {
	Save(ctx context.Context, author *Author) error
	FindByEmail(ctx context.Context, email string) (*Author, error)
}

type Author struct {
	ID          int
	Name        string
	Email       string
	Description string
	CreatedAt   time.Time
}

func NewAuthor(name, email, description string) Author {
	return Author{
		Name:        name,
		Email:       email,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

func (a Author) String() string {
	return fmt.Sprintf(
		`Author{ID: %d, Name: %s, Email: %s, Description: %s, CreatedAt: %v}`,
		a.ID,
		a.Name,
		a.Email,
		a.Description,
		a.CreatedAt,
	)
}
