package ai

import (
	"testing"

	"github.com/assertgo/assert"
)

func TestChoosesBetterMove(t *testing.T) {
	assert := assert.Setup(t)
	root := makeRoot().setChildren(
		"c1", &MockPosition{eval: 0},
		"c2", &MockPosition{eval: 3},
	)
	move := AlphaBeta(root, 1)
	assert.ThatString(move).IsEqualTo("c2")
}

func TestChoosesBetterDeep(t *testing.T) {
	assert := assert.Setup(t)
	root := makeRoot().setChildren(
		"c1", (&MockPosition{eval: 0}).setChildren(
			"c11", (&MockPosition{eval: 2}),
		), "c2", (&MockPosition{eval: 3}).setChildren(
			"c21", (&MockPosition{eval: -1}),
		),
	)
	move := AlphaBeta(root, 2)
	assert.ThatString(move).IsEqualTo("c1")
}

func TestChoosesBetterWhenWinning(t *testing.T) {
	assert := assert.Setup(t)
	root := makeRoot().setChildren(
		"c1", (&MockPosition{eval: 1000, finished: true}),
		"c2", (&MockPosition{eval: 3}).setChildren(
			"c21", (&MockPosition{eval: -1}),
		),
	)
	move := AlphaBeta(root, 2)
	assert.ThatString(move).IsEqualTo("c1")
}

func TestChoosesBetterWhenLosing(t *testing.T) {
	assert := assert.Setup(t)
	root := makeRoot().setChildren(
		"c1", (&MockPosition{eval: -1000, finished: true}),
		"c2", (&MockPosition{eval: 3}).setChildren(
			"c21", (&MockPosition{eval: -1}),
		),
	)
	move := AlphaBeta(root, 2)
	assert.ThatString(move).IsEqualTo("c2")
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
