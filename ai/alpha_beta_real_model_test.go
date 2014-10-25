package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveForwardAsFirstBestMove(t *testing.T) {
	pos := model.NewPosition()
	move := AlphaBeta(pos, 1)
	assert.Equal(t, "e2", move)
}

func TestMoveForwardAsSecondBestMove(t *testing.T) {
	pos := model.NewPosition("e2")
	move := AlphaBeta(pos, 1)
	assert.Equal(t, "e8", move)
}
