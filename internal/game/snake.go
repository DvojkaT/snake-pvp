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
