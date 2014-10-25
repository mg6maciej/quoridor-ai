package ai

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChoosesBetterMove(t *testing.T) {
	root := makeRoot()
	child1 := &MockPosition{eval: 0}
	child2 := &MockPosition{eval: 3}
	root.children = map[string]*MockPosition{"c1": child1, "c2": child2}
	move := AlphaBeta(root, 1)
	assert.Equal(t, "c2", move)
}

func TestChoosesBetterDeep(t *testing.T) {
	root := makeRoot()
	child1 := &MockPosition{eval: 0}
	child11 := &MockPosition{eval: 2}
	child1.children = map[string]*MockPosition{"c11": child11}
	child2 := &MockPosition{eval: 3}
	child21 := &MockPosition{eval: -1}
	child2.children = map[string]*MockPosition{"c21": child21}
	root.children = map[string]*MockPosition{"c1": child1, "c2": child2}
	move := AlphaBeta(root, 2)
	assert.Equal(t, "c1", move)
}

func makeRoot() *MockPosition {
	return &MockPosition{evalFunc: func() int { panic("unsupported") }, whiteActive: true}
}
