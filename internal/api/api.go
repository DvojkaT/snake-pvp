package api

import (
	"net/http"

	"github.com/centrifugal/centrifuge"
	"github.com/gin-gonic/gin"
)

func HandleRoutes(r *gin.Engine, node *centrifuge.Node) {
	websocketHandler := centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{
		CheckOrigin: func(r *http.Request) bool {
			return true //todo
		},
		ReadBufferSize:     1024,
		UseWriteBufferPool: true,
	})

	r.GET("/connection/websocket", gin.WrapH(websocketHandler))
	r.GET("/health", func(c *gin.Context) {
		OK(c, gin.H{"message": "ok"})
	})
}
