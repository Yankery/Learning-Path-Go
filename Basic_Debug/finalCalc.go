package main

func finalCalc(itemOpeningStock, itemSold, itemReturned, itemMissing int) int {
	var itemClosingStock int
	itemClosingStock = itemOpeningStock - itemSold
	itemClosingStock += itemReturned
	itemClosingStock -= itemMissing

	return itemClosingStock
}
