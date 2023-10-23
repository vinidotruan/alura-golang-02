package main

import "fmt"

type BankAccount struct {
	titular       string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	bankAccount := BankAccount{
		"Ruan",
		123,
		123,
		200.00,
	}

	fmt.Println(bankAccount)
}
