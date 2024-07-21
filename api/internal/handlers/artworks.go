package handlers

import (
	"api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArtworks(c *gin.Context) {
	// Fetch artworks
	artworks, err := services.FetchArtworks(1, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, artworks)
}
