package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed numbers.txt
	data []byte
)

func main() {
	fmt.Println(string(data))
	fmt.Println(strings.Repeat("-", 15))

	lines := strings.Split(string(data), "\r\n")

	product := 1
	for _, line := range lines {
		if line != "" {
			val, _ := strconv.Atoi(line)
			product *= val
		}
	}

	fmt.Println("Product of all numbers: ", product)
}
