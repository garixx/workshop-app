package repository

import (
	"context"
	"fmt"
	"github.com/garixx/workshop-app/internal/config"
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
