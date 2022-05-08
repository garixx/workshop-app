package inventory

import "github.com/garixx/workshop-app/internal/domain"

type Inventory struct {
	User      domain.UserUsecase
	AuthToken domain.AuthTokenUsecase
}

func NewInventory(userUsecase domain.UserUsecase, tokenUsecase domain.AuthTokenUsecase) *Inventory {
	return &Inventory{
		User:      userUsecase,
		AuthToken: tokenUsecase,
	}
}
