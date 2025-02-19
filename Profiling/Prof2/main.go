package main

import (
	"math"
	"os"
	"runtime/pprof"
	"sync"
)

func allocateMemory() {
	slice := make([]byte, 1024*1024*1024)
	for i := 0; i < len(slice); i++ {
		slice[i] = 1
	}
}

func allocateMap() {
	m := make(map[int]int, 10000000)
	for i := 0; i < 10000000; i++ {
		m[i] = i
	}
}

func processMap() {
	m := make(map[int]float64, 10000000)
	for i := 0; i < len(m); i++ {
		m[i] = float64(i)
	}
	for _, v := range m {
		v = math.Sqrt(v)
	}
}

func processSlice() {
	processMap()
	s := make([]float64, 10000000)
	for i := 0; i < len(s); i++ {
		s[i] = math.Sin(float64(i))
	}
}

func processSliceParallel() {
	processSlice()
	s := make([]float64, 10000000)
	numWorkers := 4
	numElementsPerWorker := len(s) / numWorkers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(startIndex int) {
			defer wg.Done()
			for j := startIndex; j < startIndex+numElementsPerWorker; j++ {
				s[j] = math.Sin(float64(j))
			}
		}(i * numElementsPerWorker)
	}
	wg.Wait()
}

func main() {
	f, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Call memory-intensive functions
	allocateMemory()
	allocateMap()
	processSliceParallel() // processSliceParallel -> processSlice -> processMap

	// Write memory profile to file
	//err = pprof.Lookup("heap").WriteTo(f, 0)
	err = pprof.WriteHeapProfile(f)
	if err != nil {
		panic(err)
	}
}
