package services

import (
	"Conveter/internal/models"
	"Conveter/internal/repository"
)

func ListCoins() []models.Currency {
	return repository.GetAll()
}
