package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhiteMovePawn(t *testing.T) {
	pos := NewPosition()
	pos = pos.Move("e2")
	assert.Equal(t, "e2", pos.White())
}
