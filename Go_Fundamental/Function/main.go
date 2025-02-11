package main

import (
	"fmt"
	"strings"
)

type sound struct {
	name  string
	sound string
}

func main() {
	sounds := []sound{
		{"cat", "meow"},
		{"dog", "woof"},
	}
	shouldContinue := true

	for shouldContinue {
		fmt.Println("1. Enter an animal name")
		fmt.Println("2. Print report")
		fmt.Println("q. Quit")
		fmt.Println("\nPlease select an option")

		var option string
		fmt.Scan(&option)

		switch option {
		case "1":
			sounds = append(sounds, addAnimal())
		case "2":
			printReport(sounds)
		case "q":
			shouldContinue = false
		}
	}
}

func printReport(sounds []sound) {
	fmt.Println("Animal Sounds")
	fmt.Println(strings.Repeat("-", 15))

	for _, value := range sounds {
		fmt.Println(value.name, value.sound)
	}
}

func addAnimal() sound {
	fmt.Println("Enter an animal name and it's sound")
	var newName, newSound string
	fmt.Scan(&newName, &newSound)

	return sound{name: newName, sound: newSound}
}
