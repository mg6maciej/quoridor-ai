package model

import (
	"gopkg.in/fatih/set.v0"
)

func legalMoves(pos Position) *set.Set {
	moves := set.New()
	moves.Merge(getPawnMoves(pos))
	moves.Merge(getWallsMoves(pos))
	return moves
}

func getPawnMoves(pos Position) *set.Set {
	moves := set.New()
	var pawn []rune
	if pos.WhiteActive() {
		pawn = []rune(pos.White())
	} else {
		pawn = []rune(pos.Black())
	}
	directions := [][]int32{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	sideJumps := [][][]int32{{{0, -1}, {0, 1}}, {{-1, 0}, {1, 0}}, {{-1, 0}, {1, 0}}, {{0, -1}, {0, 1}}}
	for i, dir := range directions {
		move := []rune{pawn[0] + dir[0], pawn[1] + dir[1]}
		if blockedOrOutside(pos, pawn, move) {
			continue
		}
		if string(move) == pos.White() || string(move) == pos.Black() {
			jump := []rune{move[0] + dir[0], move[1] + dir[1]}
			if blockedOrOutside(pos, move, jump) {
				for _, dir2 := range sideJumps[i] {
					sideJump := []rune{move[0] + dir2[0], move[1] + dir2[1]}
					if blockedOrOutside(pos, move, sideJump) {
						continue
					}
					moves.Add(string(sideJump))
				}
			} else {
				moves.Add(string(jump))
			}
		} else {
			moves.Add(string(move))
		}
	}
	return moves
}

func blockedOrOutside(pos Position, init []rune, dest []rune) bool {
	if dest[0] < 'a' || dest[0] > 'i' || dest[1] < '1' || dest[1] > '9' {
		return true
	}
	var wall1, wall2 string
	if init[0] == dest[0] {
		smaller := min(init[1], dest[1])
		wall1 = string([]rune{init[0], smaller, 'h'})
		wall2 = string([]rune{init[0] - 1, smaller, 'h'})
	} else if init[1] == dest[1] {
		smaller := min(init[0], dest[0])
		wall1 = string([]rune{smaller, init[1], 'v'})
		wall2 = string([]rune{smaller, init[1] - 1, 'v'})
	}
	for _, move := range pos.Walls() {
		if move == wall1 || move == wall2 {
			return true
		}
	}
	return false
}

func min(a int32, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func getWallsMoves(pos Position) *set.Set {
	moves := set.New()
	var wallsLeft int
	if pos.WhiteActive() {
		wallsLeft = pos.WhiteWallsLeft()
	} else {
		wallsLeft = pos.BlackWallsLeft()
	}
	if wallsLeft > 0 {
		for _, file := range "abcdefgh" {
			for _, rank := range "12345678" {
				for _, dir := range "hv" {
					moves.Add(string([]rune{file, rank, dir}))
				}
			}
		}
		for _, move := range pos.Walls() {
			r := []rune(move)
			r[2] = 'h'
			moves.Remove(string(r))
			r[2] = 'v'
			moves.Remove(string(r))
			r = []rune(move)
			if r[2] == 'h' {
				r[0] = rune(int(r[0]) - 1)
				moves.Remove(string(r))
				r[0] = rune(int(r[0]) + 2)
				moves.Remove(string(r))
			} else {
				r[1] = rune(int(r[1]) - 1)
				moves.Remove(string(r))
				r[1] = rune(int(r[1]) + 2)
				moves.Remove(string(r))
			}
		}
		for _, move := range set.StringSlice(moves) {
			blocking := blockingPawn(pos, move)
			if blocking {
				moves.Remove(move)
			}
		}
	}
	return moves
}

func blockingPawn(pos Position, move string) bool {
	pos = pos.Move(move)
	blocking := !legalPosition(pos)
	pos = pos.Takeback()
	return blocking
}

func legalPosition(pos Position) bool {
	return canReach(pos, pos.White(), '9') && canReach(pos, pos.Black(), '1')
}

func canReach(pos Position, start string, end rune) bool {
	return distanceToFinish(pos, start, end) >= 0
}

func distanceToFinish(pos Position, start string, end rune) int {
	if []rune(start)[1] == end {
		return 0
	}
	distance := 1
	directions := [][]int32{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	reached := set.New()
	current := set.New(start)
	next := set.New()
	for !current.IsEmpty() {
		for _, square := range set.StringSlice(current) {
			s := []rune(square)
			for _, dir := range directions {
				n := []rune{s[0] + dir[0], s[1] + dir[1]}
				if blockedOrOutside(pos, s, n) {
					continue
				}
				if s[1]+dir[1] == end {
					return distance
				}
				if !reached.Has(string(n)) {
					next.Add(string(n))
				}
			}
		}
		reached.Merge(current)
		current.Clear()
		current, next = next, current
		distance += 1
	}
	return -1
}
