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
	receivedOrderCh := receiveOrders()
	validOrderCh, invalidOrderCh := validateOrders(receivedOrderCh)
	reservedOrderCh := reserveIntenvory(validOrderCh)
	fillOrders(reservedOrderCh, &wg)

	wg.Add(1)
	go func(invalidOrderCh <-chan invalidOrder) {
		for order := range invalidOrderCh {
			fmt.Printf("Invalid order received: %vIssue: %v\n\n", order.order, order.err)
		}
		wg.Done()
	}(invalidOrderCh)

	//single producer, multiple consumer
	// const worker = 3
	// wg.Add(worker)
	// for i := 0; i < worker; i++ {
	// 	go func(reservedOrderCh <-chan order) {
	// 		for order := range reservedOrderCh {
	// 			fmt.Println("Inventory reserved for: ", order)
	// 		}
	// 		wg.Done()
	// 	}(reservedOrderCh)
	// }
	wg.Wait()
}

// multiple producer, multiple consumer
func fillOrders(in <-chan order, wg *sync.WaitGroup) {
	const worker = 3
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func() {
			for o := range in {
				o.Status = filled
				fmt.Println("Order completed: ", o)
			}
			wg.Done()
		}()
	}
}

// multiple producer, single consumer
func reserveIntenvory(in <-chan order) <-chan order {
	out := make(chan order)

	var wg sync.WaitGroup

	const worker = 3
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func() {
			for o := range in {
				o.Status = reserved
				out <- o
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func validateOrders(in <-chan order) (chan order, chan invalidOrder) {
	out := make(chan order)
	errCh := make(chan invalidOrder, 1)
	go func() {
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
	}()
	return out, errCh
}

func receiveOrders() chan order {
	out := make(chan order)
	go func() {
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
	}()
	return out
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
