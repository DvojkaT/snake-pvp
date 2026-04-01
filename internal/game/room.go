package game

import (
	"github.com/google/uuid"
)

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
	ID           string
	Snakes       map[string]Snake
	Players      map[string]LobbyPlayer
	Status       status
	Cells        [][]Cell
	PlayersLimit int64
}

func NewRoom(sizeX, sizeY, playersLimit int64) *Room {
	cells := make([][]Cell, sizeX)
	for i := range cells {
		cells[i] = make([]Cell, sizeY)
	}
	return &Room{
		ID:           uuid.NewString(),
		Snakes:       map[string]Snake{},
		Players:      map[string]LobbyPlayer{},
		Status:       LOBBY,
		Cells:        cells,
		PlayersLimit: playersLimit,
	}
}

func (r *Room) AddPlayer(ID, name string) error {
	if r.Status != LOBBY {
		return GameIsActiveError
	}

	_, ok := r.Players[ID]
	if ok {
		return UserAlreadyExistsError
	}

	if len(r.Players) >= int(r.PlayersLimit) {
		return PlayersLimitError
	}

	player := LobbyPlayer{
		ID,
		name,
	}

	r.Players[ID] = player

	return nil
}

func (r *Room) RemovePlayer(ID string) error {
	_, ok := r.Players[ID]
	if !ok {
		return UserNotFoundError
	}
	delete(r.Players, ID)
	return nil
}

func (r *Room) StartGame() error {
	var index int

	if r.Status != LOBBY {
		return GameIsActiveError
	}

	for player, _ := range r.Players {
		r.Snakes[player] = *r.setSnakePosition(index, player)
		index++
	}

	r.Status = ACTIVE

	return nil
}

// todo В будущем пересмотреть логику расставления чтобы можно было иметь 4+ игроков
func (r *Room) setSnakePosition(index int, userID string) *Snake {
	switch index {
	case 0:
		{
			snake := newSnake([]Point{}, DOWN, userID, RED)
			for p := 0; p < SnakeSize; p++ {
				point := newPoint(0, int64(p))
				r.Cells[0][p] = *NewCell(SnakePart, snake)
				snake.points = append(snake.points, *point)
			}

			return snake
		}
	case 1:
		{
			snake := newSnake([]Point{}, DOWN, userID, RED)
			lastX := len(r.Cells) - 1
			for p := 0; p < SnakeSize; p++ {
				point := newPoint(int64(lastX), int64(p))
				r.Cells[lastX][p] = *NewCell(SnakePart, snake)
				snake.points = append(snake.points, *point)
			}

			return snake
		}
	case 2:
		{
			snake := newSnake([]Point{}, UP, userID, RED)
			lastY := len(r.Cells[0]) - 1
			for p := lastY; p >= lastY-SnakeSize+1; p-- {
				point := newPoint(int64(0), int64(p))
				r.Cells[0][p] = *NewCell(SnakePart, snake)
				snake.points = append(snake.points, *point)
			}

			return snake
		}
	case 3:
		{
			{
				snake := newSnake([]Point{}, UP, userID, RED)
				lastX := len(r.Cells) - 1
				lastY := len(r.Cells[0]) - 1
				for p := lastY; p >= lastY-SnakeSize+1; p-- {
					point := newPoint(int64(lastX), int64(p))
					r.Cells[lastX][p] = *NewCell(SnakePart, snake)
					snake.points = append(snake.points, *point)
				}

				return snake
			}
		}
	}

	return nil
}
