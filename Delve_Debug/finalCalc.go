package main

func finalStockCalc(itemOpeningStock, itemSold, itemReturned int) int {
	itemClosingStock := itemOpeningStock - itemSold + itemReturned
	return itemClosingStock
}

func finalCostCalc(itemClosingStock, itemCost int) int {
	closingStockCost := itemClosingStock * itemCost
	return closingStockCost
}
