package game

type Point struct {
	x int64
	y int64
}

func newPoint(x, y int64) *Point {
	return &Point{x, y}
}
