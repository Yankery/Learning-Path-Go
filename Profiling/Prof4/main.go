package main

import (
	"flag"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

func expensiveMath() {
	for i := 0; i < 100000000; i++ {
		_ = math.Pow(2, 10)
	}
}

func allocateMemory() {
	slice := make([]byte, 1024*1024*1024)
	for i := 0; i < len(slice); i++ {
		slice[i] = 1
	}
}

func main() {
	cpuProfile := flag.String("cpuprofile", "", "write cpu profile to file")
	memProfile := flag.String("memprofile", "", "write memory profile to file")
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		defer func() {
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal(err)
			}
		}()
	}

	expensiveMath()
	allocateMemory()
}
