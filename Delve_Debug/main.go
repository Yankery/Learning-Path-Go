package main

import "fmt"

func main() {
	itemOpeningStock := 500
	itemSold := 100
	itemReturned := 50
	itemCost := 10

	itemClosingStock := finalStockCalc(itemOpeningStock, itemSold, itemReturned)
	finalCost := finalCostCalc(itemClosingStock, itemCost)

	fmt.Println("Available Inventory Count : ", itemClosingStock)
	fmt.Println("Available Inventory Cost : ", finalCost)
}
