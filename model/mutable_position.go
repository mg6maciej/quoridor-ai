package model

import (
	"gopkg.in/fatih/set.v0"
)

type MutablePosition struct {
	moves []string
}

func (pos *MutablePosition) Move(move string) Position {
	return &MutablePosition{moves: pos.movesWithMove(move)}
}

func (pos *MutablePosition) movesWithMove(move string) []string {
	moves := make([]string, len(pos.moves)+1)
	copy(moves, pos.moves)
	moves[len(moves)-1] = move
	return moves
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

func (pos *MutablePosition) Walls() []string {
	walls := make([]string, 0, len(pos.moves))
	for _, move := range pos.moves {
		if len(move) == 3 {
			walls = append(walls, move)
		}
	}
	return walls
}

func (pos *MutablePosition) Finished() bool {
	return finished(pos)
}

func (pos *MutablePosition) Eval() int {
	return eval(pos)
}

func (pos *MutablePosition) LegalMoves() *set.Set {
	return legalMoves(pos)
}
