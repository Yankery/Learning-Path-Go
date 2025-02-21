package main

import (
	"fmt"
	"sync"
	"time"
)

// not practical, since there is no need for goroutine if all go sequen
func main() {
	startNow := time.Now()
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

	// var wg sync.WaitGroup
	// const iteration = 1000
	// dataChan := make(chan int)

	// for i := 0; i < iteration; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		dataChan <- 1
	// 	}()
	// }
	// go func() {
	// 	wg.Wait()
	// 	close(dataChan)
	// }()
	// base := 0
	// for num := range dataChan {
	// 	base += num
	// }

	// fmt.Println(base)
	fmt.Println("The operation took: ", time.Since(startNow))
}
