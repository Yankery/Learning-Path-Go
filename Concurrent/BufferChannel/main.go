package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	//if channel is unbuffered, code can't be executed past line 8 because scheduler can't find receiver
	ch <- "message"

	fmt.Println(<-ch)
}
