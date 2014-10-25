package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"gopkg.in/fatih/set.v0"
)

type EvalFunc func(model.Position) int

func AlphaBeta(pos model.Position, depth int, eval EvalFunc) string {
	_, bestMove := alphaBetaImpl(pos, depth, eval, -1000000000, 1000000000)
	return bestMove
}

func alphaBetaImpl(pos model.Position, depth int, eval EvalFunc, alpha int, beta int) (int, string) {
	if depth == 0 || model.Finished(pos) {
		return eval(pos), ""
	}
	if pos.WhiteActive() {
		var bestMove string
		for _, move := range set.StringSlice(model.LegalMoves(pos)) {
			pos = pos.Move(move)
			value, _ := alphaBetaImpl(pos, depth-1, eval, alpha, beta)
			pos = pos.Takeback()
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
		for _, move := range set.StringSlice(model.LegalMoves(pos)) {
			pos = pos.Move(move)
			value, _ := alphaBetaImpl(pos, depth-1, eval, alpha, beta)
			pos = pos.Takeback()
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
