package repository

import (
	"Conveter/internal/models"
	"sync"
)

var (
	coins  = make(map[int]models.Currency)
	nextID = 1
	mu     sync.RWMutex
)

func GetAll() []models.Currency {
	mu.RLock()
	defer mu.RUnlock()

	res := make([]models.Currency, 0, len(coins))
	for _, u := range coins {
		res = append(res, u)
	}
	return res
}
