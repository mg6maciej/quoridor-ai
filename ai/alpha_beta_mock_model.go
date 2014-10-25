package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"gopkg.in/fatih/set.v0"
)

type MockPosition struct {
	eval        int
	evalFunc    func() int
	children    map[string]*MockPosition
	finished    bool
	parent      *MockPosition
	whiteActive bool
}

func (pos *MockPosition) Move(move string) model.Position {
	return pos.children[move]
}

func (pos *MockPosition) Takeback() model.Position {
	return pos.parent
}

func (pos *MockPosition) WhiteActive() bool {
	return pos.whiteActive
}

func (pos *MockPosition) Finished() bool {
	return pos.finished
}

func (pos *MockPosition) Eval() int {
	if pos.evalFunc != nil {
		return pos.evalFunc()
	}
	return pos.eval
}

func (pos *MockPosition) LegalMoves() *set.Set {
	moves := set.New()
	for move, _ := range pos.children {
		moves.Add(move)
	}
	return moves
}

func (pos *MockPosition) White() string {
	panic("unsupported")
}

func (pos *MockPosition) Black() string {
	panic("unsupported")
}

func (pos *MockPosition) WhiteWallsLeft() int {
	panic("unsupported")
}

func (pos *MockPosition) BlackWallsLeft() int {
	panic("unsupported")
}

func (pos *MockPosition) Walls() []string {
	panic("unsupported")
}
