package main

import "fmt"

type order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

type invalidOrder struct {
	order order
	err   error
}
type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

func (o order) String() string {
	return fmt.Sprintf("Product code: %v, Quantity: %v, Status: %v\n",
		o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(o orderStatus) string {
	switch o {
	case none:
		return "None"
	case new:
		return "New"
	case received:
		return "Received"
	case reserved:
		return "Reserved"
	case filled:
		return "Filled"
	default:
		return "Unknown status"
	}
}
