package usecase

import (
	"github.com/garixx/workshop-app/internal/domain"
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
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) GetUser(login string, password string) (domain.User, error) {
	return u.userRepo.GetUser(login, password)
}
