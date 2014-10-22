package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"gopkg.in/fatih/set.v0"
)

func AlphaBeta(pos *model.MutablePosition, depth int) []string {
	_, bestMoves := alphaBetaImpl(pos, depth, -1000000000, 1000000000)
	return bestMoves
}

func alphaBetaImpl(pos *model.MutablePosition, depth int, alpha int, beta int) (int, []string) {
	if depth == 0 || pos.Finished() {
		return pos.Eval(), []string{}
	}
	if pos.WhiteActive() {
		var bestMoves []string
		for _, move := range set.StringSlice(pos.LegalMoves()) {
			pos.Move(move)
			value, childMoves := alphaBetaImpl(pos, depth-1, alpha, beta)
			pos.Takeback()
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
		for _, move := range set.StringSlice(pos.LegalMoves()) {
			pos.Move(move)
			value, childMoves := alphaBetaImpl(pos, depth-1, alpha, beta)
			pos.Takeback()
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
