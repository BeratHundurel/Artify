package handlers

import (
	"api/internal/services"
	"api/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEvents(c *gin.Context) {
	page, limit := utils.ReturnUrlParameters(c)

	artworks, err := services.FetchEvents(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, artworks)
}
