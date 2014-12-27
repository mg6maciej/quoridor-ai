package model

import (
	"testing"

	"github.com/assertgo/assert"
)

func TestStartingPositionWalls(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition()
	assert.ThatInt(pos.WhiteWallsLeft()).IsEqualTo(10)
	assert.ThatInt(pos.BlackWallsLeft()).IsEqualTo(10)
}

func TestWhiteWallsLeftDecremented(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("a1v")
	assert.ThatInt(pos.WhiteWallsLeft()).IsEqualTo(9)
	assert.ThatInt(pos.BlackWallsLeft()).IsEqualTo(10)
}

func TestWhiteWallsLeftNotDecrementedOnPawnMove(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2")
	assert.ThatInt(pos.WhiteWallsLeft()).IsEqualTo(10)
	assert.ThatInt(pos.BlackWallsLeft()).IsEqualTo(10)
}

func TestBothWallsLeftDecremented(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("a1v", "b1v")
	assert.ThatInt(pos.WhiteWallsLeft()).IsEqualTo(9)
	assert.ThatInt(pos.BlackWallsLeft()).IsEqualTo(9)
}
