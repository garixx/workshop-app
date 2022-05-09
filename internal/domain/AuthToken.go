package domain

import "time"

type AuthToken struct {
	Id        int       `json:"-"`
	Login     string    `json:"login"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	ExpiredIn int       `json:"expiredIn" db:"expired_in"`
}

// AuthTokenParams user with additional data payload, e.g. expiration time in seconds
type AuthTokenParams struct {
	User     User
	ExpireIn int
	Token    AuthToken
}

type AuthTokenUsecase interface {
	StoreToken(params AuthTokenParams) (AuthToken, error)
	ValidateToken(token string) (AuthToken, error)
	GenerateToken(user User) (string, error)
	IsExpired(token AuthToken) bool
}

type AuthTokenRepository interface {
	StoreToken(authTokenParams AuthTokenParams) (AuthToken, error)
	FetchToken(token string) (AuthToken, error)
}
