package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/garixx/workshop-app/internal/domain"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

const tokensTable = "tokens"

var InvalidTokenError = errors.New("invalid token")

type PostgresAuthTokenRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresAuthTokenRepository(pool *pgxpool.Pool) *PostgresAuthTokenRepository {
	return &PostgresAuthTokenRepository{pool: pool}
}

func (p PostgresAuthTokenRepository) StoreToken(params domain.AuthTokenParams) (domain.AuthToken, error) {
	var newToken domain.AuthToken
	query := fmt.Sprintf("INSERT INTO %s (login, token, expired_in) VALUES ($1, $2, $3) RETURNING id, login, token, created_at, expired_in", tokensTable)
	row := p.pool.QueryRow(context.Background(), query, params.User.Login, params.Token.Token, params.ExpireIn)
	if err := row.Scan(&newToken.Id, &newToken.Login, &newToken.Token, &newToken.CreatedAt, &newToken.ExpiredIn); err != nil {
		return domain.AuthToken{}, err
	}

	return newToken, nil
}

func (p PostgresAuthTokenRepository) FetchToken(token string) (domain.AuthToken, error) {
	var tokens []*domain.AuthToken
	err := pgxscan.Select(context.Background(), p.pool, &tokens, "select * from tokens where token = $1", token)
	if err != nil {
		return domain.AuthToken{}, err
	}
	if len(tokens) != 1 {
		return domain.AuthToken{}, InvalidTokenError
	}

	return *tokens[0], nil
}
