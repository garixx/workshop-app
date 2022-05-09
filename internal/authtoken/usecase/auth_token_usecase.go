package usecase

import (
	"errors"
	"github.com/garixx/workshop-app/internal/domain"
	"time"
)

type AuthTokenUsecase struct {
	userRepo domain.AuthTokenRepository
}

func NewAuthTokenUsecase(repo domain.AuthTokenRepository) domain.AuthTokenUsecase {
	return &AuthTokenUsecase{
		userRepo: repo,
	}
}

func (a AuthTokenUsecase) GenerateToken(user domain.User) (string, error) {
	token := user.Login + "xxxxx"
	return token, nil
}

func (a AuthTokenUsecase) IsExpired(token domain.AuthToken) bool {
	if time.Now().After(token.CreatedAt.Add(time.Duration(token.ExpiredIn) * time.Second)) {
		return true
	}
	return false
}

func (a AuthTokenUsecase) StoreToken(params domain.AuthTokenParams) (domain.AuthToken, error) {
	// check defaults
	if params.ExpireIn < 0 {
		params.ExpireIn = 43200
	}
	generated, err := a.GenerateToken(params.User)
	if err != nil {
		return domain.AuthToken{}, errors.New("token generation failed")
	}
	params.Token = domain.AuthToken{Token: generated}
	return a.userRepo.StoreToken(params)
}

func (a AuthTokenUsecase) ValidateToken(token string) (domain.AuthToken, error) {
	return a.userRepo.FetchToken(token)
}
