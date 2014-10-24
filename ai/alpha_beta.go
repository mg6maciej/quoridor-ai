package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"gopkg.in/fatih/set.v0"
)

func AlphaBeta(pos model.Position, depth int) []string {
	_, bestMoves := alphaBetaImpl(pos, depth, -1000000000, 1000000000)
	return bestMoves
}

func alphaBetaImpl(pos model.Position, depth int, alpha int, beta int) (int, []string) {
	if depth == 0 || model.Finished(pos) {
		return pos.Eval(), []string{}
	}
	if pos.WhiteActive() {
		var bestMoves []string
		for _, move := range set.StringSlice(model.LegalMoves(pos)) {
			pos = pos.Move(move)
			value, childMoves := alphaBetaImpl(pos, depth-1, alpha, beta)
			pos = pos.Takeback()
			if alpha < value {
				alpha = value
				bestMoves = []string{move}
				bestMoves = append(bestMoves, childMoves...)
			}
			if beta <= alpha {
				break
			}
		}
		return alpha, bestMoves
	} else {
		var bestMoves []string
		for _, move := range set.StringSlice(model.LegalMoves(pos)) {
			pos = pos.Move(move)
			value, childMoves := alphaBetaImpl(pos, depth-1, alpha, beta)
			pos = pos.Takeback()
			if beta > value {
				beta = value
				bestMoves = []string{move}
				bestMoves = append(bestMoves, childMoves...)
			}
			if beta <= alpha {
				break
			}
		}
		return beta, bestMoves
	}
}
