package book

type CreateBookRequest struct {
	Title          string  `json:"title" binding:"required"`
	Abstract       string  `json:"abstract" binding:"required,max=500"`
	TableOfContent string  `json:"table_of_content" binding:"required"`
	Price          float32 `json:"price" binding:"required,min=20"`
	NumberOfPages  int     `json:"number_of_pages" binding:"required,min=100"`
	ISBN           string  `json:"isbn" binding:"required"`
	PublishDate    string  `json:"publish_date" binding:"required"`
	Category       string  `json:"category" binding:"required"`
	AuthorID       int     `json:"author_id" binding:"required"`
}

type BookDTO struct {
	ID             int     `json:"id"`
	Title          string  `json:"title"`
	Abstract       string  `json:"abstract"`
	TableOfContent string  `json:"table_of_content"`
	Price          float32 `json:"price"`
	NumberOfPages  int     `json:"number_of_pages"`
	ISBN           string  `json:"isbn"`
	PublishDate    string  `json:"publish_date"`
	Category       string  `json:"category"`
	AuthorID       int     `json:"author_id"`
}
