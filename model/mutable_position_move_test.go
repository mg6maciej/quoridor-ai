package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhiteMovePawn(t *testing.T) {
	pos := NewPosition()
	pos = pos.Move("e2")
	assert.Equal(t, "e2", pos.White())
}

func TestTakeback(t *testing.T) {
	pos := NewPosition("e2")
	pos = pos.Takeback()
	assert.Equal(t, "e1", pos.White())
}
