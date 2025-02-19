package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"sync"

	_ "net/http/pprof"
)

func expensiveRand() {
	for i := 0; i < 1000000000; i++ {
		b := make([]byte, 16)
		_, _ = rand.Read(b)
	}
}

func allocateMemory() {
	slice := make([]byte, 1024*1024*1024)
	for i := 0; i < len(slice); i++ {
		slice[i] = 1
	}
}

func main() {
	var wg sync.WaitGroup

	// Server for pprof
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	wg.Add(2) // increment WaitGroup counter by 2

	// Run expensiveMath and signal to WaitGroup when done
	go func() {
		defer wg.Done()
		expensiveRand()
	}()

	// Run expensiveRand and signal to WaitGroup when done
	go func() {
		defer wg.Done()
		allocateMemory()
	}()

	// Wait for both goroutines to complete
	wg.Wait()
}
