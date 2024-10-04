package user

import "time"

type User struct {
	ID        int
	Login     string
	Password  string
	CreatedAt time.Time
}

func NewUser(login, password string) User {
	return User{
		Login:     login,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
