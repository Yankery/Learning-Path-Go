package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receiveOrderCh = make(chan order)
	var validOrderCh = make(chan order)
	var invalidOrderCh = make(chan invalidOrder)
	go receiveOrders(receiveOrderCh)
	go validateOrders(receiveOrderCh, validOrderCh, invalidOrderCh)

	wg.Add(1)
	go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
		select {
		case order := <-validOrderCh:
			fmt.Println("Valid order received: ", order)
		case order := <-invalidOrderCh:
			fmt.Println("Invalid order received: ", order.order)
			fmt.Println("Issue: ", order.err)
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}

func validateOrders(in <-chan order, out chan<- order, errCh chan<- invalidOrder) {
	order := <-in
	if order.Quantity <= 0 {
		errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
	} else {
		out <- order
	}
}

func receiveOrders(out chan<- order) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		out <- newOrder
	}
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
