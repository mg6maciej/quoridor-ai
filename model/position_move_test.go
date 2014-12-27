package model

import (
	"testing"

	"github.com/assertgo/assert"
)

func TestWhiteMovePawn(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition()
	pos = pos.Move("e2")
	assert.ThatString(pos.White()).IsEqualTo("e2")
}
