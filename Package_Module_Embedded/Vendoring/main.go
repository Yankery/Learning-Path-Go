package main

import (
	"calc"
	"fmt"

	"github.com/pioz/faker"
)

var itemPrice = 50
var itemDiscount int

func main() {
	itemDiscount = 75
	Curr := faker.CurrencySymbol()
	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", Curr, itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", Curr, totalDiscount)
	fmt.Println("\nRandom Color: ", faker.ColorName())
}
