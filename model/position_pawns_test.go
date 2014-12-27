package model

import (
	"testing"

	"github.com/assertgo/assert"
)

func TestStartingPositionPawns(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition()
	assert.ThatString(pos.White()).IsEqualTo("e1")
	assert.ThatString(pos.Black()).IsEqualTo("e9")
}

func TestStartingPositionActive(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition()
	assert.ThatBool(pos.WhiteActive()).IsTrue()
}

func TestPositionAfterOneMove(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2")
	assert.ThatString(pos.White()).IsEqualTo("e2")
	assert.ThatBool(pos.WhiteActive()).IsFalse()
}

func TestPositionAfterWallMove(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e5h")
	assert.ThatString(pos.White()).IsEqualTo("e1")
}

func TestPositionAfterTwoMoves(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2", "e8")
	assert.ThatString(pos.White()).IsEqualTo("e2")
	assert.ThatString(pos.Black()).IsEqualTo("e8")
}

func TestPositionAfterThreeMoves(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2", "e8", "e3")
	assert.ThatString(pos.White()).IsEqualTo("e3")
	assert.ThatString(pos.Black()).IsEqualTo("e8")
}
