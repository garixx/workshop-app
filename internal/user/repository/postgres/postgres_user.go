package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/garixx/workshop-app/internal/config"
	"github.com/garixx/workshop-app/internal/domain"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

const usersTable = "users"

func NewPostgresDB(cfg config.PostgresDBConfig) (*pgxpool.Pool, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	ctx := context.Background()

	db, err := pgxpool.Connect(ctx, conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository(pool *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{pool: pool}
}

//
//func (a *PostgresUserRepository) CreateUser(user domain.User) (domain.User, error) {
//	var id int
//	var login string
//	var username string
//	var createdAt time.Time
//	query := fmt.Sprintf("INSERT INTO %s (login, username, password_hash) VALUES ($1, $2, $3) RETURNING id, login, username, created_at", usersTable)
//	row := a.pool.QueryRow(context.Background(), query, user.Login, user.Username, user.PasswordHash)
//	if err := row.Scan(&id, &login, &username, &createdAt); err != nil {
//		return domain.User{}, err
//	}
//
//	return domain.User{
//		Id:        id,
//		Login:     login,
//		Username:  username,
//		CreatedAt: createdAt,
//	}, nil
//}

func (a *PostgresUserRepository) CreateUser(user domain.User) (domain.User, error) {
	var newUser domain.User
	query := fmt.Sprintf("INSERT INTO %s (login, username, password_hash) VALUES ($1, $2, $3) RETURNING id, login, username, created_at", usersTable)
	row := a.pool.QueryRow(context.Background(), query, user.Login, user.Username, user.PasswordHash)
	if err := row.Scan(&newUser.Id, &newUser.Login, &newUser.Username, &newUser.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}

func (a *PostgresUserRepository) GetUser(username string, password string) (domain.User, error) {
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
