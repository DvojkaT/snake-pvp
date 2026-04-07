package api

import (
	"fmt"
	"net/http"
	"snake/internal/game"

	"github.com/centrifugal/centrifuge"
	"github.com/gin-gonic/gin"
)

func HandleRoutes(r *gin.Engine, node *centrifuge.Node, roomsList game.RoomList) {
	websocketHandler := centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{
		CheckOrigin: func(r *http.Request) bool {
			return true //todo
		},
		ReadBufferSize:     1024,
		UseWriteBufferPool: true,
	})

	r.GET("/connection/websocket", gin.WrapH(websocketHandler))
	r.POST("/:gameId/join", func(c *gin.Context) {
		joinGame(c, roomsList)
	})
	r.POST("/:gameId/start", func(c *gin.Context) {
		startGame(c, roomsList)
	})
	r.GET("/health", func(c *gin.Context) {
		OK(c, gin.H{"message": "ok"})
	})
}

func joinGame(c *gin.Context, roomsList game.RoomList) {
	joinGameRequest := struct {
		UserId string `json:"user_id" binding:"required"`
		Name   string `json:"name" binding:"required"`
	}{}
	gameId := c.Param("gameId")
	if err := c.ShouldBindJSON(&joinGameRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	room, ok := roomsList[gameId]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
	}
	player := game.NewLobbyPlayer(joinGameRequest.UserId, "Test player")
	room.Players[player.ID] = *player
	OK(c, gin.H{"message": "ok"}) //todo Добавить вывод цвета
}

func startGame(c *gin.Context, roomsList game.RoomList) {
	gameId := c.Param("gameId")
	room, ok := roomsList[gameId]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
	}
	err := room.StartGame()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not start game"})
	}
	room.StartTicker()
	OK(c, gin.H{"message": "ok"}) //todo Добавить вывод цвета //
}
