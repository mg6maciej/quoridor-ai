package model

import (
	"testing"

	"github.com/assertgo/assert"
	"gopkg.in/fatih/set.v0"
)

func TestInitialLegalMoves(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition()
	moves := pos.LegalMoves()
	expected := set.New("e2", "d1", "f1")
	addAllWalls(expected)
	assert.That(moves).IsEqualTo(expected)
}

func TestInitialBlackLegalMoves(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2")
	moves := pos.LegalMoves()
	expected := set.New("e8", "d9", "f9")
	addAllWalls(expected)
	assert.That(moves).IsEqualTo(expected)
}

func TestLegalMovesAfterWall(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e5h")
	moves := pos.LegalMoves()
	expected := set.New("e8", "d9", "f9")
	addAllWalls(expected)
	expected.Remove("e5h", "e5v", "d5h", "f5h")
	assert.That(moves).IsEqualTo(expected)
}

func TestLegalMovesAfterAllWalls(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("a1v", "d9", "b1v", "e9", "c1v", "d9", "f1v", "e9", "g1v", "d9", "h1v", "e9", "a3v", "d9", "b3v", "e9", "c3v", "d9", "d3v", "e9")
	moves := pos.LegalMoves()
	expected := set.New("e2", "d1", "f1")
	assert.That(moves).IsEqualTo(expected)
}

func TestHasStraightJumpMove(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5")
	moves := pos.LegalMoves()
	expected := set.New("e4", "e7", "d6", "f6")
	addAllWalls(expected)
	assert.That(moves).IsEqualTo(expected)
}

func TestCannotJumpOverHorizontalWall(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e8h")
	moves := pos.LegalMoves()
	expected := set.New("d9", "f9")
	addAllWalls(expected)
	expected.Remove("d8h", "e8h", "e8v", "f8h")
	assert.That(moves).IsEqualTo(expected)
}

func TestCannotJumpOverVerticalWall(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e8v")
	moves := pos.LegalMoves()
	expected := set.New("d9", "e8")
	addAllWalls(expected)
	expected.Remove("e8h", "e7v", "e8v")
	assert.That(moves).IsEqualTo(expected)
}

func TestHasSideJumpMove(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5", "d6h")
	moves := pos.LegalMoves()
	expected := set.New("e4", "d5", "f5", "d6", "f6")
	addAllWalls(expected)
	expected.Remove("c6h", "d6h", "d6v", "e6h")
	assert.That(moves).IsEqualTo(expected)
}

func TestBlockedSideJumpMove(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5", "e4h", "d5v")
	moves := pos.LegalMoves()
	expected := set.New("e7", "f6", "f5")
	addAllWalls(expected)
	expected.Remove("d4h", "e4h", "e4v", "f4h", "d4v", "d5h", "d5v", "d6v")
	assert.That(moves).IsEqualTo(expected)
}

func TestCannotCompletlyBlockWhite(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("e1h", "f1v")
	moves := pos.LegalMoves()
	assert.ThatBool(moves.Has("d1v")).IsFalse()
}

func TestCannotCompletlyBlockBlack(t *testing.T) {
	assert := assert.Setup(t)
	pos := NewPosition("d8v", "e8v")
	moves := pos.LegalMoves()
	assert.ThatBool(moves.Has("d7h")).IsFalse()
	assert.ThatBool(moves.Has("e7h")).IsFalse()
}

func addAllWalls(s *set.Set) {
	for _, file := range "abcdefgh" {
		for _, rank := range "12345678" {
			for _, dir := range "hv" {
				s.Add(string([]rune{file, rank, dir}))
			}
		}
	}
}
