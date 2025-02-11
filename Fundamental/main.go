package main

import (
	"fmt"
	"strings"
)

func main() {
	//Simple data types
	name, age := "David", 12
	employee := []string{"Arthur", "David", "Carter"}
	ages := []int{11, 12, 13}

	fmt.Println(name, ": ", age)
	fmt.Println(strings.Repeat("-", 15))
	for i := 0; i < len(employee); i++ {
		fmt.Println(employee[i], ": ", ages[i])
	}
}
