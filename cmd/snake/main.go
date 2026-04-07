package main

import (
	"fmt"
	"log"
	"snake/internal/api"
	"snake/internal/game"
	"snake/internal/ws"
	"time"

	"github.com/centrifugal/centrifuge"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}
	corsConfig.AllowAllOrigins = true
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

	api.HandleRoutes(r, node, list)
	ws.HandleConnection(node, list)

	// Заглушка
	room := game.NewRoom(list, 50, 50, 4)

	go func() {
		for {
			data := <-room.ViewState
			err := ws.PublishRoomState(node, data)
			if err != nil {
				fmt.Printf("Error publishing room state: %v\n", err)
			}
		}
	}()

	// Вот по сюда

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
