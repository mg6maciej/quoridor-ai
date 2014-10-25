package model

import (
	"gopkg.in/fatih/set.v0"
)

type ImmutablePosition struct {
	parent         *ImmutablePosition
	white          string
	black          string
	whiteActive    bool
	whiteWallsLeft int
	blackWallsLeft int
	walls          []string
}

func (pos *ImmutablePosition) Move(move string) Position {
	child := &ImmutablePosition{
		parent:         pos,
		white:          pos.whiteAfterMove(move),
		black:          pos.blackAfterMove(move),
		whiteActive:    !pos.whiteActive,
		whiteWallsLeft: pos.whiteWallsLeftAfterMove(move),
		blackWallsLeft: pos.blackWallsLeftAfterMove(move),
		walls:          pos.wallsAfterMove(move),
	}
	return child
}

func (pos *ImmutablePosition) Takeback() Position {
	return pos.parent
}

func (pos *ImmutablePosition) White() string {
	return pos.white
}

func (pos *ImmutablePosition) Black() string {
	return pos.black
}

func (pos *ImmutablePosition) WhiteActive() bool {
	return pos.whiteActive
}

func (pos *ImmutablePosition) WhiteWallsLeft() int {
	return pos.whiteWallsLeft
}

func (pos *ImmutablePosition) BlackWallsLeft() int {
	return pos.blackWallsLeft
}

func (pos *ImmutablePosition) Walls() []string {
	return pos.walls
}

func (pos *ImmutablePosition) Finished() bool {
	return Finished(pos)
}

func (pos *ImmutablePosition) Eval() int {
	return Eval(pos)
}

func (pos *ImmutablePosition) LegalMoves() *set.Set {
	return LegalMoves(pos)
}

func (pos *ImmutablePosition) whiteAfterMove(move string) string {
	if pos.whiteActive && len(move) == 2 {
		return move
	}
	return pos.white
}

func (pos *ImmutablePosition) blackAfterMove(move string) string {
	if !pos.whiteActive && len(move) == 2 {
		return move
	}
	return pos.black
}

func (pos *ImmutablePosition) whiteWallsLeftAfterMove(move string) int {
	wallsLeft := pos.whiteWallsLeft
	if pos.whiteActive && len(move) == 3 {
		wallsLeft--
	}
	return wallsLeft
}

func (pos *ImmutablePosition) blackWallsLeftAfterMove(move string) int {
	wallsLeft := pos.blackWallsLeft
	if !pos.whiteActive && len(move) == 3 {
		wallsLeft--
	}
	return wallsLeft
}

func (pos *ImmutablePosition) wallsAfterMove(move string) []string {
	if len(move) == 3 {
		walls := make([]string, 0, len(pos.walls)+1)
		walls = append(walls, pos.walls...)
		walls = append(walls, move)
		return walls
	}
	return pos.walls
}
