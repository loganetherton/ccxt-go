package main

import (
	"github.com/loganetheton/ccxt-go/cmd"
	"runtime"
)

func main() {
	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd.Execute()
}
