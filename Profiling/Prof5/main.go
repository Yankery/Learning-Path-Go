package main

import (
	"crypto/rand"
	"math"
	"os"
	"runtime/pprof"
)

func expensiveMath() {
	for i := 0; i < 100000000; i++ {
		_ = math.Pow(2, 10)
	}
}

func expensiveRand() {
	for i := 0; i < 100000000; i++ {
		b := make([]byte, 16)
		_, _ = rand.Read(b)
	}
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// Expensive functions
	expensiveMath()
	expensiveRand()
}
