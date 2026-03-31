package api

import "github.com/gin-gonic/gin"

func HandleRoutes(g *gin.Engine) {
	g.GET("/health", func(c *gin.Context) {
		OK(c, gin.H{"message": "ok"})
	})
}
