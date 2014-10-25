package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"gopkg.in/fatih/set.v0"
)

func AlphaBeta(pos model.Position, depth int) string {
	_, bestMove := alphaBetaImpl(pos, depth, -1000000000, 1000000000)
	return bestMove
}

func alphaBetaImpl(pos model.Position, depth int, alpha int, beta int) (int, string) {
	if depth == 0 || pos.Finished() {
		return pos.Eval(), ""
	}
	if pos.WhiteActive() {
		var bestMove string
		for _, move := range set.StringSlice(pos.LegalMoves()) {
			child := pos.Move(move)
			value, _ := alphaBetaImpl(child, depth-1, alpha, beta)
			if alpha < value {
				alpha = value
				bestMove = move
			}
			if beta <= alpha {
				break
			}
		}
		return alpha, bestMove
	} else {
		var bestMove string
		for _, move := range set.StringSlice(pos.LegalMoves()) {
			child := pos.Move(move)
			value, _ := alphaBetaImpl(child, depth-1, alpha, beta)
			if beta > value {
				beta = value
				bestMove = move
			}
			if beta <= alpha {
				break
			}
		}
		return beta, bestMove
	}
}
