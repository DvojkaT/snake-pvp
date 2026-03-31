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
