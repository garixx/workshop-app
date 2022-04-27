package repository

import (
	"github.com/garixx/workshop-app/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Authorization interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(username string, password string) (domain.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(pool),
	}
}
