package ai

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChoosesBetterMove(t *testing.T) {
	root := makeRoot().setChildren(
		"c1", &MockPosition{eval: 0},
		"c2", &MockPosition{eval: 3},
	)
	move := AlphaBeta(root, 1)
	assert.Equal(t, "c2", move)
}

func TestChoosesBetterDeep(t *testing.T) {
	root := makeRoot().setChildren(
		"c1", (&MockPosition{eval: 0}).setChildren(
			"c11", (&MockPosition{eval: 2}),
		), "c2", (&MockPosition{eval: 3}).setChildren(
			"c21", (&MockPosition{eval: -1}),
		),
	)
	move := AlphaBeta(root, 2)
	assert.Equal(t, "c1", move)
}

func makeRoot() *MockPosition {
	return &MockPosition{evalFunc: func() int { panic("unsupported") }, whiteActive: true}
}

func (parent *MockPosition) setChildren(children ...interface{}) *MockPosition {
	parent.children = make(map[string]*MockPosition)
	for i := 0; i < len(children); i += 2 {
		parent.children[children[i].(string)] = children[i+1].(*MockPosition)
	}
	return parent
}
