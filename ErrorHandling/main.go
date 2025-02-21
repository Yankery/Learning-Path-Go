package main

import "fmt"

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	fmt.Println("Starting risky operation")
	panic("Something went horribly wrong!")
	fmt.Println("This will not be executed")
}

func main() {
	riskyOperation()
	fmt.Println("Program continues after recovery")
}
