package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment usong Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using Bank Account %d\n", bankAccount)
}

//adaptador que implementa la estructura de BankPayment
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

//función que implementa el Pay() para el tipo BankPaymentAdapter que posee el Pay(int)
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	//error porque el BankPayment no tiene la implementación de Pay()
	//bank := &BankPayment{}
	//ProcessPayment(bank)

	bpa := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		bankAccount: 20,
	}
	ProcessPayment(bpa)
}
