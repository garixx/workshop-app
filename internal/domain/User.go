package domain

import "time"

type User struct {
	Id        int       `json:"id"`
	Login     string    `json:"login"`
	Username  string    `json:"username"`
	Password  string    `json:"password" db:"password_hash"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type UserUsecase interface {
	CreateUser(user User) (User, error)
	GetUser(user User) (User, error)
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUser(user User) (User, error)
}
