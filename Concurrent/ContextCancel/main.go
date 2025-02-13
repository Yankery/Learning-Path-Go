package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background()) //parent context
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for range time.Tick(500 * time.Millisecond) {
			//stop for loop if context is canceled
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			fmt.Println("Tick")
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	//cancel the context
	cancel()
	wg.Wait()
}
