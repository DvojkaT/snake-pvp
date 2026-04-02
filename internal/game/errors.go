package game

import "errors"

var (
	GameIsActiveError      = errors.New("game is active")
	UserAlreadyExistsError = errors.New("user already exists")
	UserNotFoundError      = errors.New("user not found")
	PlayersLimitError      = errors.New("players limit")
	DirectionNotFoundError = errors.New("direction not found")
	CellTypeNotFoundError  = errors.New("cell type not found")
)
