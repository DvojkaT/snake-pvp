package game

import "github.com/google/uuid"

type status int64

const (
	LOBBY status = iota
	ACTIVE
	END
)

type LobbyPlayer struct {
	ID   string
	Name string
}

type Room struct {
	ID      string
	Snakes  map[string]Snake
	Players map[string]LobbyPlayer
	Status  status
	Cells   [][]Cell
}

func NewRoom(sizeX, sizeY int64) *Room {
	cells := make([][]Cell, sizeX)
	for i := range cells {
		cells[i] = make([]Cell, sizeY)
	}
	return &Room{
		ID:      uuid.NewString(),
		Snakes:  map[string]Snake{},
		Players: map[string]LobbyPlayer{},
		Status:  LOBBY,
		Cells:   cells,
	}
}
