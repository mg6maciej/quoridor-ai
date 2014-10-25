package ai

import (
	"github.com/mg6maciej/quoridor-ai/model"
	"gopkg.in/fatih/set.v0"
)

type EvalFunc func(model.Position) int

type LegalMovesFunc func(model.Position) *set.Set

type FinishedFunc func(model.Position) bool

func AlphaBeta(pos model.Position, depth int, eval EvalFunc, legalMoves LegalMovesFunc, finished FinishedFunc) string {
	_, bestMove := alphaBetaImpl(pos, depth, eval, legalMoves, finished, -1000000000, 1000000000)
	return bestMove
}

func alphaBetaImpl(pos model.Position, depth int, eval EvalFunc, legalMoves LegalMovesFunc, finished FinishedFunc, alpha int, beta int) (int, string) {
	if depth == 0 || finished(pos) {
		return eval(pos), ""
	}
	if pos.WhiteActive() {
		var bestMove string
		for _, move := range set.StringSlice(legalMoves(pos)) {
			pos = pos.Move(move)
			value, _ := alphaBetaImpl(pos, depth-1, eval, legalMoves, finished, alpha, beta)
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
		for _, move := range set.StringSlice(legalMoves(pos)) {
			pos = pos.Move(move)
			value, _ := alphaBetaImpl(pos, depth-1, eval, legalMoves, finished, alpha, beta)
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
