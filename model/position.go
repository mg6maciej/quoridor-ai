package model

type Position interface {
	Move(move string) Position
	Takeback() Position
	White() string
	Black() string
	WhiteActive() bool
	WhiteWallsLeft() int
	BlackWallsLeft() int
	Walls() []string
}

func NewPosition(moves ...string) Position {
	return &MutablePosition{moves}
}

func Finished(pos Position) bool {
	return []rune(pos.White())[1] == '9' || []rune(pos.Black())[1] == '1'
}
