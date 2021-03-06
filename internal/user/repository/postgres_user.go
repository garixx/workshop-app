package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/garixx/workshop-app/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

const usersTable = "users"

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository(pool *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{pool: pool}
}

func (a *PostgresUserRepository) CreateUser(user models.User) (models.User, error) {
	var newUser models.User
	query := fmt.Sprintf("INSERT INTO %s (login, username, password_hash) VALUES ($1, $2, $3) RETURNING id, login, username, created_at", usersTable)
	row := a.pool.QueryRow(context.Background(), query, user.Login, user.Username, user.Password)
	if err := row.Scan(&newUser.Id, &newUser.Login, &newUser.Username, &newUser.CreatedAt); err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

func (a *PostgresUserRepository) GetUser(user models.User) (models.User, error) {
	var users []*models.User
	err := pgxscan.Select(context.Background(), a.pool, &users, "select login from users where login = $1 and password_hash = $2", user.Login, user.Password)
	if err != nil {
		return models.User{}, err
	}
	if len(users) != 1 {
		return models.User{}, errors.New(fmt.Sprintf("incorrect number of rows. expected one but got %d", len(users)))
	}

	return *users[0], nil
}
