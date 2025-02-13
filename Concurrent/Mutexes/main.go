package main

import (
	"fmt"
	"sync"
)

// not practical, since there is no need for goroutine if all go sequen
func main() {
	s := []int{}
	var wg sync.WaitGroup
	var m sync.Mutex
	const iteration = 1000

	wg.Add(iteration)
	for i := 0; i < iteration; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			s = append(s, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(len(s))
}
