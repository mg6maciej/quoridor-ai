package ai

import (
	"testing"

	"github.com/assertgo/assert"
	"github.com/mg6maciej/quoridor-ai/model"
)

func TestMoveForwardAsFirstBestMove(t *testing.T) {
	assert := assert.Setup(t)
	pos := model.NewPosition()
	move := AlphaBeta(pos, 1)
	assert.ThatString(move).IsEqualTo("e2")
}

func TestMoveForwardAsSecondBestMove(t *testing.T) {
	assert := assert.Setup(t)
	pos := model.NewPosition("e2")
	move := AlphaBeta(pos, 1)
	assert.ThatString(move).IsEqualTo("e8")
}
