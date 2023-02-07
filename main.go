package main

import (
	"github.com/loganetherton/ccxt-go/cmd"
	"runtime"
)

func main() {
	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd.Execute()
}
