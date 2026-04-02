package game

type Direction int64

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

const SnakeSize = 3

type Color string

const (
	RED    Color = "#FF0000"
	BLUE   Color = "#2E27F5"
	YELLOW Color = "#F2F527"
	GREEN  Color = "#00FF04"
	PINK   Color = "#EA00FF"
	CYAN   Color = "#00FFFF"
	BLACK  Color = "#000000"
)

type Snake struct {
	points    []Point
	direction Direction
	userID    string
	color     Color
}

func newSnake(points []Point, direction Direction, userID string, color Color) *Snake {
	return &Snake{
		points:    points,
		direction: direction,
		userID:    userID,
		color:     color,
	}
}

// Move Обработка движения змейки. В случае если возвщрается хвост, значит змейка укоротилась на его размер
func (s *Snake) Move(ateFruit bool, sizeX, sizeY int64) (head *Point, tail *Point, err error) {
	head, err = s.NextHead(sizeX, sizeY)
	if err != nil {
		return nil, nil, err
	}
	s.points = append(s.points, *head)
	if !ateFruit {
		tailPoint := s.points[0]
		tail = &tailPoint
		s.points = s.points[1:]
	}
	return head, tail, nil
}

// NextHead Получение следующего места головы змейки
func (s *Snake) NextHead(sizeX, sizeY int64) (*Point, error) {
	var dx, dy int64

	switch s.direction {
	case UP:
		{
			dy = -1
		}
	case DOWN:
		{
			dy = 1
		}
	case RIGHT:
		{
			dx = 1
		}
	case LEFT:
		{
			dx = -1
		}
	default:
		return nil, DirectionNotFoundError
	}

	currentHead := &s.points[len(s.points)-1]
	newX := (currentHead.x + dx + sizeX) % sizeX
	newY := (currentHead.y + dy + sizeY) % sizeY

	return newPoint(newX, newY), nil
}
