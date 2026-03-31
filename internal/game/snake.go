package game

import (
	"snake/internal/direction"
)

type Direction int64

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Snake struct {
	points    []Point
	direction Direction
	userId    string
}
