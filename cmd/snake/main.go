package main

import (
	"embed"
	"log"
	"os"
	"snake/internal/api"
	"snake/internal/game"
	"snake/internal/ws"
	"strings"
	"time"

	"github.com/centrifugal/centrifuge"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed static
var FS embed.FS

func main() {
	r := gin.Default()

	corsOrigin := os.Getenv("CORS_ALLOW_ORIGIN")
	corsConfig := cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}
	if corsOrigin == "*" || corsOrigin == "" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = strings.Split(corsOrigin, ",")
		corsConfig.AllowCredentials = true
	}
	r.Use(cors.New(corsConfig))

	node, err := centrifuge.New(centrifuge.Config{
		LogLevel:       centrifuge.LogLevelInfo,
		LogHandler:     handleLog,
		HistoryMetaTTL: 24 * time.Hour,
	})
	if err != nil {
		log.Fatal(err)
	}

	list := game.NewRoomList()

	api.HandleRoutes(r, node, &FS, list)
	ws.HandleConnection(node, list)

	go func() {
		err = node.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

func handleLog(e centrifuge.LogEntry) {
	log.Printf("%s: %v", e.Message, e.Fields)
}
