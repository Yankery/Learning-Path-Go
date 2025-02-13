package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{}
	var wg sync.WaitGroup
	const iteration = 1000

	wg.Add(iteration)
	for i := 0; i < iteration; i++ {
		go func() {
			s = append(s, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(len(s))
}
