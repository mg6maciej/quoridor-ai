package main

import (
	"flag"
	"fmt"
	"github.com/mg6maciej/quoridor-ai/ai"
	"github.com/mg6maciej/quoridor-ai/model"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	pos := model.NewPosition()
	for i := 1; i <= 3; i++ {
		start := time.Now()
		move := ai.AlphaBeta(pos, i, model.Eval, model.LegalMoves, model.Finished)
		elapsed := time.Since(start)
		fmt.Printf("%v %s\n", move, elapsed)
	}
}
