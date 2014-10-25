package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartingPositionEvaluatesToZero(t *testing.T) {
	pos := NewPosition()
	eval := pos.Eval()
	assert.Equal(t, 0, eval)
}

func TestWonPositionEvaluatesToBigPositiveValue(t *testing.T) {
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5", "a1v", "e7", "e8", "e9")
	eval := pos.Eval()
	assert.InDelta(t, 1000000, eval, 100)
}

func TestLostPositionEvaluatesToBigNegativeValue(t *testing.T) {
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5", "e4", "e3", "e2", "e1", "d1")
	eval := pos.Eval()
	assert.InDelta(t, -1000000, eval, 100)
}

func TestAfterMoveEvaluatesToPositiveValue(t *testing.T) {
	pos := NewPosition("e2")
	eval := pos.Eval()
	assert.Equal(t, 1, eval)
}

func TestAfterWallEvaluatesToNegativeValue(t *testing.T) {
	pos := NewPosition("a1v")
	eval := pos.Eval()
	assert.InDelta(t, -2, eval, 1)
}
