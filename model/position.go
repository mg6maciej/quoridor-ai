package model

import (
	"gopkg.in/fatih/set.v0"
)

type Position interface {
	Finished() bool
	Move(move string) Position
	Takeback() Position
	White() string
	Black() string
	WhiteActive() bool
	WhiteWallsLeft() int
	BlackWallsLeft() int
	Eval() int
	LegalMoves() *set.Set
}

func NewPosition(moves ...string) Position {
	return &MutablePosition{moves}
}
