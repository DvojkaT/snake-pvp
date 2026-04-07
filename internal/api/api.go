package api

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"snake/internal/game"
	"snake/internal/ws"

	"github.com/centrifugal/centrifuge"
	"github.com/gin-gonic/gin"
)

func HandleRoutes(r *gin.Engine, node *centrifuge.Node, FS *embed.FS, roomsList game.RoomList) {
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
	r.POST("/create", func(c *gin.Context) {
		createRoom(c, node, roomsList)
	})
	r.GET("/health", func(c *gin.Context) {
		OK(c, gin.H{"message": "ok"})
	})

	sub, err := fs.Sub(FS, "static/assets")
	if err != nil {
		log.Fatal(err)
	}
	r.StaticFS("/assets", http.FS(sub))

	r.NoRoute(func(c *gin.Context) {
		data, err := FS.ReadFile("static/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
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
		return
	}
	room, ok := roomsList[gameId]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
		return
	}
	_ = room.AddPlayer(joinGameRequest.UserId, joinGameRequest.Name)
	OK(c, gin.H{"message": "ok"}) //todo Добавить вывод цвета
}

func startGame(c *gin.Context, roomsList game.RoomList) {
	gameId := c.Param("gameId")
	room, ok := roomsList[gameId]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
		return
	}
	err := room.StartGame()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not start game"})
		return
	}
	room.StartTicker()
	OK(c, gin.H{"message": "ok"})
}

func createRoom(c *gin.Context, node *centrifuge.Node, roomsList game.RoomList) {
	room := game.NewRoom(roomsList, 50, 50, 4)

	go func() {
		for {
			select {
			case data := <-room.ViewState:
				{
					err := ws.PublishRoomState(node, data)
					if err != nil {
						fmt.Printf("Error publishing room state: %v\n", err)
					}
				}
			case data := <-room.LobbyState:
				{
					err := ws.PublishLobbyState(node, room.ID, data)
					if err != nil {
						fmt.Printf("Error publishing room state: %v\n", err)
					}
				}
			case <-room.StopTimer:
				return
			}
		}
	}()
	OK(c, gin.H{"uuid": room.ID})
}
