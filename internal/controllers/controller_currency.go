package controllers

import (
	"Conveter/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCoins(c *gin.Context) {
	coins := services.ListCoins()
	c.JSON(http.StatusOK, gin.H{"data": coins})
}
