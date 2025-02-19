package main

import (
	"fmt"
	"log"
	"oop/payment"
)

type PaymentProcessor[T payment.Float] interface {
	ProcessPayment(amount T) error
}

type Account[T payment.Float] interface {
	Available() T
}

type PaymentMethod[T payment.Float] interface {
	PaymentProcessor[T]
	Account[T]
}

func main() {
	var pm PaymentMethod[float64] = payment.NewCreditCard[float64]("Debra", "1111-2222", 12, 1999, 123, 5000)
	pm = payment.NewBankAccount[float64]("Debra", "1234", 4000)
	err := pm.ProcessPayment(10000)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Processed payment. Remaining credit: ", pm.Available())
	}

	err = pm.ProcessPayment(1000)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Processed payment. Remaining credit: ", pm.Available())
	}

	switch m := pm.(type) {
	case *payment.CreditCard[float64]:
		fmt.Printf("Credit Card: %T\n", m)
	case *payment.BankAccount[float64]:
		fmt.Printf("Bank Account: %T\n", m)
	}

}
