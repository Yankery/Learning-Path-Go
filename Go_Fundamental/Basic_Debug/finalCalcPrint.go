package main

import "fmt"

func finalCalcPrint(itemOpeningStock, itemSold, itemReturned, itemMissing int) int {
	var itemClosingStock int
	fmt.Println("ItemOpeningStock is: ", itemOpeningStock)
	fmt.Println("ItemClosingStock is: ", itemClosingStock)
	itemClosingStock = itemOpeningStock - itemSold
	fmt.Println("itemSold is: ", itemSold)
	fmt.Println("ItemClosingStock is: ", itemClosingStock)
	itemClosingStock += itemReturned
	fmt.Println("itemReturned is: ", itemReturned)
	fmt.Println("ItemClosingStock is: ", itemClosingStock)
	itemClosingStock += itemMissing
	fmt.Println("itemMissing is: ", itemMissing)
	fmt.Println("ItemClosingStock is: ", itemClosingStock)

	return itemClosingStock
}
