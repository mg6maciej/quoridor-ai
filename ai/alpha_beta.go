package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
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
		pos.LegalMoves().Each(func(item interface{}) bool {
			move := item.(string)
			child := pos.Move(move)
			value, _ := alphaBetaImpl(child, depth-1, alpha, beta)
			if alpha < value {
				alpha = value
				bestMove = move
			}
			return !(beta <= alpha)
		})
		return alpha, bestMove
	} else {
		var bestMove string
		pos.LegalMoves().Each(func(item interface{}) bool {
			move := item.(string)
			child := pos.Move(move)
			value, _ := alphaBetaImpl(child, depth-1, alpha, beta)
			if beta > value {
				beta = value
				bestMove = move
			}
			return !(beta <= alpha)
		})
		return beta, bestMove
	}
}
