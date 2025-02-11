package main

import (
	"fmt"
)

func main() {
	// name, age := "David", 12

	// employee := []string{"Arthur", "David", "Carter"}
	// ages := []int{21, 22, 23}

	// employeeMap := map[string]int{
	// 	employee[0]: ages[0],
	// 	employee[1]: ages[1],
	// 	employee[2]: ages[2],
	// }

	type form struct {
		name string
		age  int
	}

	employeeStruct := []form{
		{"Arthur", 21},
		{"David", 22},
		{"Carter", 23},
	}

	// fmt.Println(name, ": ", age)
	// fmt.Println(strings.Repeat("-", 15))

	// fmt.Print("Please type employee name: ")
	// var option string
	// fmt.Scan(&option)

	//If statement
	// if option == "Arthur" {
	// 	fmt.Println("Arthur", ": ", employeeMap["Arthur"])
	// } else if option == "David" {
	// 	fmt.Println("David", ": ", employeeMap["David"])
	// } else if option == "Carter" {
	// 	fmt.Println("Carter", ": ", employeeMap["Carter"])
	// } else {
	// 	fmt.Println("Employee doesn't exist!")
	// }

	//Switch statement
	// val, ok := employeeMap[option]
	// switch ok {
	// case true:
	// 	fmt.Println(option, ": ", val)
	// default:
	// 	fmt.Println("Employee doesn't exist!")
	// }

	//infinite for loop
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	//full conditional for loop
	// for i := 0; i < len(employee); i++ {
	// 	fmt.Println(employee[i], ": ", ages[i])
	// }
	// fmt.Println(strings.Repeat("-", 15))

	//collection-based for loop
	for _, value := range employeeStruct {
		fmt.Println(value.name, ": ", value.age)
	}
}
