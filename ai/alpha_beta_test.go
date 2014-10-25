package ai

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChoosesBetterMove(t *testing.T) {
	root := &MockPosition{evalFunc: func() int { panic("unsupported") }, whiteActive: true}
	child1 := &MockPosition{eval: 0}
	child2 := &MockPosition{eval: 1}
	root.children = map[string]*MockPosition{"c1": child1, "c2": child2}
	move := AlphaBeta(root, 1)
	assert.Equal(t, "c2", move)
}
