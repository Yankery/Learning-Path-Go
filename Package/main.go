package main

import (
	"calc"
	"fmt"
)

var itemPrice = 50
var itemDiscount int

func init() {
	itemDiscount = 20
	fmt.Println("")
	fmt.Println(">>main.go init() itemDiscount: ", itemDiscount)
	fmt.Println("")
}

func main() {
	itemDiscount = 75
	//can't use discount, must use Discount
	//totalDiscount := calc.discount(itemPrice, itemDiscount)
	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
