package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveForwardAsFirstBestMove(t *testing.T) {
	pos := model.NewPosition()
	moves := AlphaBeta(pos, 1, true)
	assert.Equal(t, 1, len(moves))
	assert.Equal(t, "e2", moves[0])
}

func TestMoveForwardAsSecondBestMove(t *testing.T) {
	pos := model.NewPosition("e2")
	moves := AlphaBeta(pos, 1, false)
	assert.Equal(t, 1, len(moves))
	assert.Equal(t, "e8", moves[0])
}
