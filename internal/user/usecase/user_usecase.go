package usecase

import (
	"github.com/garixx/workshop-app/internal/hashgenerator"
	"github.com/garixx/workshop-app/internal/models"
)

type UserUsecase struct {
	userRepo models.UserRepository
}

func NewUserUsecase(repo models.UserRepository) models.UserRepository {
	return &UserUsecase{
		userRepo: repo,
	}
}

func (u *UserUsecase) CreateUser(user models.User) (models.User, error) {
	user.Password = hashgenerator.GeneratePasswordHash(user.Password)
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) GetUser(user models.User) (models.User, error) {
	passwordHash := hashgenerator.GeneratePasswordHash(user.Password)
	user.Password = passwordHash
	return u.userRepo.GetUser(user)
}
