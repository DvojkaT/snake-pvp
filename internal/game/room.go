package game

import (
	"fmt"
	"log"
	"time"
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
	Snakes       map[string]*Snake
	Players      map[string]LobbyPlayer
	Status       status
	Cells        [][]Cell
	PlayersLimit int64
	stopTimer    chan bool
	ViewState    chan *RoomView
}

type RoomView struct {
	ID    string       `json:"id"`
	Cells [][]CellView `json:"cells"`
}

func NewRoomView(room *Room) *RoomView {
	cells := make([][]CellView, len(room.Cells))
	for i := range cells {
		cells[i] = make([]CellView, len(room.Cells[0]))
	}

	for indexX := range room.Cells {
		for indexY, cell := range room.Cells[indexX] {
			snakeView := &SnakeView{}
			if cell.snake != nil {
				snakeView.Color = cell.snake.color
			} else {
				snakeView = nil
			}
			cells[indexX][indexY] = CellView{
				Object: cell.object,
				Snake:  snakeView,
			}
		}
	}

	return &RoomView{
		ID:    room.ID,
		Cells: cells,
	}
}

func NewLobbyPlayer(uuid, name string) *LobbyPlayer {
	return &LobbyPlayer{
		ID:   uuid,
		Name: name,
	}
}

func NewRoom(sizeX, sizeY, playersLimit int64) *Room {
	cells := make([][]Cell, sizeX)
	for i := range cells {
		cells[i] = make([]Cell, sizeY)
	}
	return &Room{
		ID:           "test-game-id", //todo uuid.NewString()
		Snakes:       map[string]*Snake{},
		Players:      map[string]LobbyPlayer{},
		Status:       LOBBY,
		Cells:        cells,
		PlayersLimit: playersLimit,
		stopTimer:    make(chan bool),
		ViewState:    make(chan *RoomView, 1),
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
		r.Snakes[player] = r.setSnakePosition(index, player)
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

func (r *Room) StartTicker() {
	ticker := time.NewTicker(time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-r.stopTimer:
				return
			case <-ticker.C:
				err := r.nextTick()
				fmt.Printf("game %s ticked\n", r.ID)
				if err != nil {
					log.Printf("tick error: %v", err)
					return
				}
			}
		}
	}()
}

func (r *Room) StopTicker() {
	close(r.stopTimer)
}

// nextTick Обработка тика каждый n времени
func (r *Room) nextTick() error {
	sizeX := len(r.Cells)
	sizeY := len(r.Cells[0])
	for _, snake := range r.Snakes {
		withFruit := false
		snakeNextHead, err := snake.NextHead(int64(sizeX), int64(sizeY))
		if err != nil {
			return err
		}
		switch r.Cells[snakeNextHead.x][snakeNextHead.y].object {
		case Fruit:
			withFruit = true
		case SnakePart:
			r.snakeLose(snake)
			return nil
		case Empty:
		default:
			return CellTypeNotFoundError
		}

		head, tail, err := snake.Move(withFruit, int64(sizeX), int64(sizeY))
		if err != nil {
			return err
		}
		fmt.Printf("snake %s. new head posotion - x: %d, y: %d\n", snake.userID, head.x, head.y)

		// Двигаем голову на новое место
		r.Cells[head.x][head.y].object = SnakePart
		r.Cells[head.x][head.y].snake = snake

		// В случае получения координат бывшего хвоста, удаляем точку
		if tail != nil {
			r.Cells[tail.x][tail.y].object = Empty
			r.Cells[tail.x][tail.y].snake = nil
		}
		r.ViewState <- NewRoomView(r)
	}
	return nil
}

func (r *Room) snakeLose(snake *Snake) {
	//todo
}
