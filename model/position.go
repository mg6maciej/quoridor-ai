package model

import (
	"gopkg.in/fatih/set.v0"
)

type Position interface {
	Move(move string) Position
	Takeback() Position
	White() string
	Black() string
	WhiteActive() bool
	WhiteWallsLeft() int
	BlackWallsLeft() int
	Walls() []string
	Finished() bool
	Eval() int
	LegalMoves() *set.Set
}

//func NewPosition(moves ...string) Position {
//	return &MutablePosition{moves}
//}

func NewPosition(moves ...string) Position {
	var pos Position = &ImmutablePosition{white: "e1", black: "e9", whiteActive: true, whiteWallsLeft: 10, blackWallsLeft: 10, walls: []string{}}
	for _, move := range moves {
		pos = pos.Move(move)
	}
	return pos
}

func finished(pos Position) bool {
	return []rune(pos.White())[1] == '9' || []rune(pos.Black())[1] == '1'
}
