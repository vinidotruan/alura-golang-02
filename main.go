package main

import (
	"bank/accounts"
	"fmt"
)
import "bank/owners"

type checkingAccount interface {
	Withdraw(value float64) string
}

func PayBill(account checkingAccount, value float64) {
	account.Withdraw(value)
}

func main() {
	owner := owners.Owner{Name: "John", Age: 25}
	bankAccount := accounts.CheckingAccount{Owner: owner, AgencyNumber: 589, AccountNumber: 123456}
	bankAccount.Deposit(100)
	fmt.Println(bankAccount.GetBalance())
	PayBill(&bankAccount, 60)
	fmt.Println(bankAccount.GetBalance())
}
