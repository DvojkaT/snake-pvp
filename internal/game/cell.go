package game

type CellType int64

const (
	Empty CellType = iota
	Fruit
	SnakePart
)

type Cell struct {
	object CellType
	snake  *Snake
}

type CellView struct {
	Object CellType   `json:"object"`
	Snake  *SnakeView `json:"snake"`
}

func NewCell(object CellType, snake *Snake) *Cell {
	return &Cell{object, snake}
}
