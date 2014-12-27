package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartingPositionPawns(t *testing.T) {
	pos := NewPosition()
	assert.Equal(t, "e1", pos.White())
	assert.Equal(t, "e9", pos.Black())
}

func TestStartingPositionActive(t *testing.T) {
	pos := NewPosition()
	assert.True(t, pos.WhiteActive())
}

func TestPositionAfterOneMove(t *testing.T) {
	pos := NewPosition("e2")
	assert.Equal(t, "e2", pos.White())
	assert.False(t, pos.WhiteActive())
}

func TestPositionAfterWallMove(t *testing.T) {
	pos := NewPosition("e5h")
	assert.Equal(t, "e1", pos.White())
}

func TestPositionAfterTwoMoves(t *testing.T) {
	pos := NewPosition("e2", "e8")
	assert.Equal(t, "e2", pos.White())
	assert.Equal(t, "e8", pos.Black())
}

func TestPositionAfterThreeMoves(t *testing.T) {
	pos := NewPosition("e2", "e8", "e3")
	assert.Equal(t, "e3", pos.White())
	assert.Equal(t, "e8", pos.Black())
}
