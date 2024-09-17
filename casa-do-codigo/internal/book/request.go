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
