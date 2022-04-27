package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/garixx/workshop-app/internal/domain"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type AuthPostgres struct {
	pool *pgxpool.Pool
}

func NewAuthPostgres(pool *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{pool: pool}
}

func (a *AuthPostgres) CreateUser(user domain.User) (domain.User, error) {
	var id int
	var login string
	var username string
	var createdAt time.Time
	query := fmt.Sprintf("INSERT INTO %s (login, username, password_hash) VALUES ($1, $2, $3) RETURNING id, login, username, created_at", usersTable)
	row := a.pool.QueryRow(context.Background(), query, user.Login, user.Username, user.PasswordHash)
	if err := row.Scan(&id, &login, &username, &createdAt); err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Id:        id,
		Login:     login,
		Username:  username,
		CreatedAt: createdAt,
	}, nil
}

func (a *AuthPostgres) GetUser(username string, password string) (domain.User, error) {
	var users []*domain.User
	err := pgxscan.Select(context.Background(), a.pool, &users, "select * from users where login = $1 and password_hash = $2", username, password)
	if err != nil {
		return domain.User{}, err
	}
	if len(users) != 1 {
		return domain.User{}, errors.New(fmt.Sprintf("incorrect number of rows. expected one but got %d", len(users)))
	}

	return *users[0], nil
}
