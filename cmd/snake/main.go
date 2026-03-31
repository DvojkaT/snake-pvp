package main

import (
	"snake/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api.HandleRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
