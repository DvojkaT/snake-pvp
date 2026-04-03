package main

import (
	"fmt"
	"log"
	"snake/internal/api"
	"snake/internal/game"
	"snake/internal/ws"
	"time"

	"github.com/centrifugal/centrifuge"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	node, err := centrifuge.New(centrifuge.Config{
		LogLevel:       centrifuge.LogLevelInfo,
		LogHandler:     handleLog,
		HistoryMetaTTL: 24 * time.Hour,
	})
	if err != nil {
		log.Fatal(err)
	}

	api.HandleRoutes(r, node)
	ws.HandleConnection(node)

	// Заглушка
	room := game.NewRoom(50, 50, 4)
	player := game.NewLobbyPlayer("testPlayerId", "Test player")
	room.Players[player.ID] = *player

	err = room.StartGame()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("game %s started\n", room.ID)
	room.StartTicker()

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
