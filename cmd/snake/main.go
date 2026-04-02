package main

import (
	"log"
	"snake/internal/api"
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

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	}()

	err = node.Run()
	if err != nil {
		return
	}
}

func handleLog(e centrifuge.LogEntry) {
	log.Printf("%s: %v", e.Message, e.Fields)
}
