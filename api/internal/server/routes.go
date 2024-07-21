package server

import (
	"api/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/artworks", handlers.GetArtworks)

	return r
}
