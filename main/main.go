package main

import (
	"fmt"
	"github.com/mg6maciej/quoridor-ai/ai"
	"github.com/mg6maciej/quoridor-ai/model"
	"time"
)

func main() {
	pos := model.NewPosition()
	for i := 1; i <= 10; i++ {
		start := time.Now()
		moves := ai.AlphaBeta(pos, i)
		elapsed := time.Since(start)
		fmt.Printf("%v %s\n", moves, elapsed)
	}
}
