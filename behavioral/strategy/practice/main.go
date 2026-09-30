package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount int) error
}

type CardPayment struct {
	CardNumber string
	CVV        string
}

func (cp *CardPayment) Pay(amount int) error {
	fmt.Printf("pay by card: %s - %s\n", cp.CardNumber, cp.CVV)
	return nil
}

type SBPPayment struct{}

func (sbp *SBPPayment) Pay(amount int) error {
	fmt.Println("pay by sbp")
	return nil
}

type Order struct {
	Amount   int
	Strategy PaymentStrategy
}

func (o *Order) SetStrategy(strategy PaymentStrategy) {
	o.Strategy = strategy
}

func (o *Order) ExecutePayment() error {
	return o.Strategy.Pay(o.Amount)
}

func main() {
	order := &Order{Amount: 100}
	order.SetStrategy(&CardPayment{CardNumber: "4242-1111-2222-3333", CVV: "123"})
	err := order.ExecutePayment()
	if err != nil {
		fmt.Println(err)
	}
	order.SetStrategy(&SBPPayment{})
	err = order.ExecutePayment()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("payment executed")

}
