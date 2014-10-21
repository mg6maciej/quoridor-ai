package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartingPositionWalls(t *testing.T) {
	pos := NewPosition()
	assert.Equal(t, 10, pos.WhiteWallsLeft())
	assert.Equal(t, 10, pos.BlackWallsLeft())
}

func TestWhiteWallsLeftDecremented(t *testing.T) {
	pos := NewPosition("a1v")
	assert.Equal(t, 9, pos.WhiteWallsLeft())
	assert.Equal(t, 10, pos.BlackWallsLeft())
}

func TestWhiteWallsLeftNotDecrementedOnPawnMove(t *testing.T) {
	pos := NewPosition("e2")
	assert.Equal(t, 10, pos.WhiteWallsLeft())
	assert.Equal(t, 10, pos.BlackWallsLeft())
}

func TestBothWallsLeftDecremented(t *testing.T) {
	pos := NewPosition("a1v", "b1v")
	assert.Equal(t, 9, pos.WhiteWallsLeft())
	assert.Equal(t, 9, pos.BlackWallsLeft())
}
