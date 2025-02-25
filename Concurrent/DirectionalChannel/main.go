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
	loop:
		for {
			select {
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Println("Valid order received: ", order)
				} else {
					break loop
				}

			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Println("Invalid order received: ", order.order)
					fmt.Println("Issue: ", order.err)
				} else {
					break loop
				}

			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}

func validateOrders(in <-chan order, out chan<- order, errCh chan<- invalidOrder) {
	//order := <-in
	for order := range in {
		if order.Quantity <= 0 {
			errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
			fmt.Println("")
		} else {
			out <- order
		}
	}
	close(out)
	close(errCh)
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
	close(out)
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": -19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
