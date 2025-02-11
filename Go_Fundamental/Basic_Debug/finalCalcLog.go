package main

import (
	"log"
	"os"
)

func finalCalcLog(itemOpeningStock, itemSold, itemReturned, itemMissing int) int {
	file, er := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		log.Fatal(er)
	}
	log.SetOutput(file)

	var itemClosingStock int
	log.Println("ItemOpeningStock is: ", itemOpeningStock)
	log.Println("ItemClosingStock is: ", itemClosingStock)
	itemClosingStock = itemOpeningStock - itemSold
	log.Println("itemSold is: ", itemSold)
	log.Println("ItemClosingStock is: ", itemClosingStock)
	itemClosingStock += itemReturned
	log.Println("itemReturned is: ", itemReturned)
	log.Println("ItemClosingStock is: ", itemClosingStock)
	itemClosingStock += itemMissing
	log.Println("itemMissing is: ", itemMissing)
	log.Println("ItemClosingStock is: ", itemClosingStock)

	return itemClosingStock
}
