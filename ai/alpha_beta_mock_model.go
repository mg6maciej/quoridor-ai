package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"gopkg.in/fatih/set.v0"
)

func MockEval(pos model.Position) int {
	mock := pos.(*MockPosition)
	if mock.evalFunc != nil {
		return mock.evalFunc()
	}
	return mock.eval
}

func MockLegalMoves(pos model.Position) *set.Set {
	moves := set.New()
	for move, _ := range pos.(*MockPosition).children {
		moves.Add(move)
	}
	return moves
}

func MockFinished(pos model.Position) bool {
	return pos.(*MockPosition).finished
}

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
