package author

import "fmt"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateAuthor(name, email, description string) {
	fmt.Println(name, email, description)
}
