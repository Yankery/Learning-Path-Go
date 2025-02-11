package main

import (
	"fmt"
	"strings"
)

func main() {
	itemOpeningStock := 500
	itemSold := 100
	itemReturned := 50
	itemMissing := 5
	var itemClosingStockCalc int

	itemClosingStockActual := 445
	fmt.Println("Available Inventory (Check): ", itemClosingStockActual)
	fmt.Println(strings.Repeat("-", 15))

	itemClosingStockCalc = finalCalc(itemOpeningStock, itemSold, itemReturned, itemMissing)
	fmt.Println("Available Inventory (Calc): ", itemClosingStockCalc)
	fmt.Println(strings.Repeat("-", 15))

	if itemClosingStockCalc != itemClosingStockActual {
		fmt.Println("Warning: Calculation ERROR!")
	} else {
		fmt.Println("Correct calculation.")
	}
}
