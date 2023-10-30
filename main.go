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

func (account *BankAccount) transfer(value float64, destinationAccount *BankAccount) bool {
	if value < account.balance && value > 0 {
		account.balance -= value
		destinationAccount.balance += value
		return true
	}
	return false
}

func main() {
	bankAccount := BankAccount{
		"Ruan",
		123,
		123,
		200.00,
	}

	bankAccount2 := BankAccount{
		"Ruan 2",
		001,
		564,
		300.,
	}

	fmt.Println(bankAccount.withdraw(100))
	fmt.Println(bankAccount.balance)
	fmt.Println(bankAccount.deposit(100))
	fmt.Println(bankAccount2.transfer(100, &bankAccount))
}
