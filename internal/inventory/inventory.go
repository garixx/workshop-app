package inventory

import "github.com/garixx/workshop-app/internal/models"

type Inventory struct {
	User      models.UserUsecase
	AuthToken models.AuthTokenUsecase
}

func NewInventory(userUsecase models.UserUsecase, tokenUsecase models.AuthTokenUsecase) *Inventory {
	return &Inventory{
		User:      userUsecase,
		AuthToken: tokenUsecase,
	}
}
