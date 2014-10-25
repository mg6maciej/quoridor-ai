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
