package model

type MutablePosition struct {
	moves []string
}

func NewPosition(moves ...string) *MutablePosition {
	return &MutablePosition{moves}
}

func (pos *MutablePosition) Finished() bool {
	return []rune(pos.White())[1] == '9' || []rune(pos.Black())[1] == '1'
}

func (pos *MutablePosition) Move(move string) {
	pos.moves = append(pos.moves, move)
}

func (pos *MutablePosition) Takeback() {
	pos.moves = pos.moves[:len(pos.moves)-1]
}

func (pos *MutablePosition) White() string {
	l := len(pos.moves)
	for i := l - 1 - (l+1)%2; i >= 0; i -= 2 {
		if len(pos.moves[i]) == 2 {
			return pos.moves[i]
		}
	}
	return "e1"
}

func (pos *MutablePosition) Black() string {
	l := len(pos.moves)
	for i := l - 1 - l%2; i >= 0; i -= 2 {
		if len(pos.moves[i]) == 2 {
			return pos.moves[i]
		}
	}
	return "e9"
}

func (pos *MutablePosition) WhiteActive() bool {
	return len(pos.moves)%2 == 0
}

func (pos *MutablePosition) WhiteWallsLeft() int {
	wallsLeft := 10
	for i, move := range pos.moves {
		if i%2 == 0 && len(move) == 3 {
			wallsLeft -= 1
		}
	}
	return wallsLeft
}

func (pos *MutablePosition) BlackWallsLeft() int {
	wallsLeft := 10
	for i, move := range pos.moves {
		if i%2 == 1 && len(move) == 3 {
			wallsLeft -= 1
		}
	}
	return wallsLeft
}
