package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func blockingDemo1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
	}()
	wg.Wait()
}

func blockingDemo2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-time.After(3 * time.Second)
	}()
	wg.Wait()
}

func blockingDemo3() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000000; i++ {
			_ = i * i
		}
	}()
	wg.Wait()
}

func blockingDemo4() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var m sync.Mutex
		m.Lock()
		time.Sleep(3 * time.Second)
		m.Unlock()
	}()
	wg.Wait()
}

func main() {
	// Create a file to write the profile to
	f, err := os.Create("block.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Enable blocking profile
	runtime.SetBlockProfileRate(1)

	blockingDemo1()
	blockingDemo2()
	blockingDemo3()
	blockingDemo4()

	// Write blocking profile to file
	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		fmt.Println(err)
		return
	}
}
