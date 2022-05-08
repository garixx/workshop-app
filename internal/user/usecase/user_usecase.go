package usecase

import (
	"github.com/garixx/workshop-app/internal/domain"
	"github.com/garixx/workshop-app/internal/helper"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserRepository {
	return &UserUsecase{
		userRepo: repo,
	}
}

func (u *UserUsecase) CreateUser(user domain.User) (domain.User, error) {
	user.Password = helper.GeneratePasswordHash(user.Password)
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) GetUser(user domain.User) (domain.User, error) {
	passwordHash := helper.GeneratePasswordHash(user.Password)
	user.Password = passwordHash
	return u.userRepo.GetUser(user)
}
