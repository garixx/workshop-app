package domain

import "time"

type User struct {
	Id           int       `json:"-"`
	Login        string    `json:"name"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password" db:"password_hash"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

type NilUser User
