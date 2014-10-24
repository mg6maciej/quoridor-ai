package model

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/fatih/set.v0"
	"testing"
)

func TestInitialLegalMoves(t *testing.T) {
	pos := NewPosition()
	moves := LegalMoves(pos)
	expected := set.New("e2", "d1", "f1")
	addAllWalls(expected)
	assert.Equal(t, expected, moves)
}

func TestInitialBlackLegalMoves(t *testing.T) {
	pos := NewPosition("e2")
	moves := LegalMoves(pos)
	expected := set.New("e8", "d9", "f9")
	addAllWalls(expected)
	assert.Equal(t, expected, moves)
}

func TestLegalMovesAfterWall(t *testing.T) {
	pos := NewPosition("e5h")
	moves := LegalMoves(pos)
	expected := set.New("e8", "d9", "f9")
	addAllWalls(expected)
	expected.Remove("e5h", "e5v", "d5h", "f5h")
	assert.Equal(t, expected, moves)
}

func TestLegalMovesAfterAllWalls(t *testing.T) {
	pos := NewPosition("a1v", "d9", "b1v", "e9", "c1v", "d9", "f1v", "e9", "g1v", "d9", "h1v", "e9", "a3v", "d9", "b3v", "e9", "c3v", "d9", "d3v", "e9")
	moves := LegalMoves(pos)
	expected := set.New("e2", "d1", "f1")
	assert.Equal(t, expected, moves)
}

func TestHasStraightJumpMove(t *testing.T) {
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5")
	moves := LegalMoves(pos)
	expected := set.New("e4", "e7", "d6", "f6")
	addAllWalls(expected)
	assert.Equal(t, expected, moves)
}

func TestCannotJumpOverHorizontalWall(t *testing.T) {
	pos := NewPosition("e8h")
	moves := LegalMoves(pos)
	expected := set.New("d9", "f9")
	addAllWalls(expected)
	expected.Remove("d8h", "e8h", "e8v", "f8h")
	assert.Equal(t, expected, moves)
}

func TestCannotJumpOverVerticalWall(t *testing.T) {
	pos := NewPosition("e8v")
	moves := LegalMoves(pos)
	expected := set.New("d9", "e8")
	addAllWalls(expected)
	expected.Remove("e8h", "e7v", "e8v")
	assert.Equal(t, expected, moves)
}

func TestHasSideJumpMove(t *testing.T) {
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5", "d6h")
	moves := LegalMoves(pos)
	expected := set.New("e4", "d5", "f5", "d6", "f6")
	addAllWalls(expected)
	expected.Remove("c6h", "d6h", "d6v", "e6h")
	assert.Equal(t, expected, moves)
}

func TestBlockedSideJumpMove(t *testing.T) {
	pos := NewPosition("e2", "e8", "e3", "e7", "e4", "e6", "e5", "e4h", "d5v")
	moves := LegalMoves(pos)
	expected := set.New("e7", "f6", "f5")
	addAllWalls(expected)
	expected.Remove("d4h", "e4h", "e4v", "f4h", "d4v", "d5h", "d5v", "d6v")
	assert.Equal(t, expected, moves)
}

func TestCannotCompletlyBlockWhite(t *testing.T) {
	pos := NewPosition("e1h", "f1v")
	moves := LegalMoves(pos)
	assert.False(t, moves.Has("d1v"))
}

func TestCannotCompletlyBlockBlack(t *testing.T) {
	pos := NewPosition("d8v", "e8v")
	moves := LegalMoves(pos)
	assert.False(t, moves.Has("d7h"))
	assert.False(t, moves.Has("e7h"))
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
