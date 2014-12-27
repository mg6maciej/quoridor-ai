package model

import (
	"testing"

	"github.com/assertgo/assert"
)

func TestStartingPositionEvaluatesToZero(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition()
	eval := pos.Eval()
	assert.ThatInt(eval).IsEqualTo(0)
}

func TestWonPositionEvaluatesToBigPositiveValue(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5", "a1v", "e7", "e8", "e9")
	eval := pos.Eval()
	assert.ThatInt(eval).IsBetween(1000000-100, 1000000+100)
}

func TestLostPositionEvaluatesToBigNegativeValue(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5", "e4", "e3", "e2", "e1", "d1")
	eval := pos.Eval()
	assert.ThatInt(eval).IsBetween(-1000000-100, -1000000+100)
}

func TestAfterMoveEvaluatesToPositiveValue(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2")
	eval := pos.Eval()
	assert.ThatInt(eval).IsEqualTo(1)
}

func TestAfterWallEvaluatesToNegativeValue(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("a1v")
	eval := pos.Eval()
	assert.ThatInt(eval).IsBetween(-2-1, -2+1)
}
