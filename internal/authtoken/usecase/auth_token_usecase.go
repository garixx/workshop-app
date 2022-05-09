package usecase

import (
	"errors"
	"github.com/garixx/workshop-app/internal/models"
	"time"
)

type AuthTokenUsecase struct {
	userRepo models.AuthTokenRepository
}

func NewAuthTokenUsecase(repo models.AuthTokenRepository) models.AuthTokenUsecase {
	return &AuthTokenUsecase{
		userRepo: repo,
	}
}

func (a AuthTokenUsecase) GenerateToken(user models.User) (string, error) {
	token := user.Login + "xxxxx"
	return token, nil
}

func (a AuthTokenUsecase) IsExpired(token models.AuthToken) bool {
	if time.Now().After(token.CreatedAt.Add(time.Duration(token.ExpiredIn) * time.Second)) {
		return true
	}
	return false
}

func (a AuthTokenUsecase) StoreToken(params models.AuthTokenParams) (models.AuthToken, error) {
	// check defaults
	if params.ExpireIn < 0 {
		params.ExpireIn = 43200
	}
	generated, err := a.GenerateToken(params.User)
	if err != nil {
		return models.AuthToken{}, errors.New("token generation failed")
	}
	params.Token = models.AuthToken{Token: generated}
	return a.userRepo.StoreToken(params)
}

func (a AuthTokenUsecase) ValidateToken(token string) (models.AuthToken, error) {
	return a.userRepo.FetchToken(token)
}
