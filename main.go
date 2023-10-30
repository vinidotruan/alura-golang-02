package main

import "fmt"

type BankAccount struct {
	titular       string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (account *BankAccount) withdraw(value float64) string {
	canDraw := value > 0 && value <= account.balance

	if canDraw {
		account.balance -= value
		return "success"
	}
	return "insufficient funds"
}

func (account *BankAccount) deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		account.balance += value
		return "success", account.balance
	}
	return "invalid value", account.balance

}

func main() {
	bankAccount := BankAccount{
		"Ruan",
		123,
		123,
		200.00,
	}

	fmt.Println(bankAccount.withdraw(100))
	fmt.Println(bankAccount.balance)
	fmt.Println(bankAccount.deposit(100))
}
