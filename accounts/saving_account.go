package accounts

import "bank/owners"

type SavingAccount struct {
	Owner         owners.Owner
	AgencyNumber  int
	AccountNumber int
	Operation     int
	balance       float64
}

func (account *SavingAccount) Withdraw(value float64) string {
	canDraw := value > 0 && value <= account.balance

	if canDraw {
		account.balance -= value
		return "success"
	}
	return "insufficient funds"
}

func (account *SavingAccount) Deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		account.balance += value
		return "success", account.balance
	}
	return "invalid value", account.balance

}

func (account *SavingAccount) Transfer(value float64, destinationAccount *SavingAccount) bool {
	if value < account.balance && value > 0 {
		account.balance -= value
		destinationAccount.balance += value
		return true
	}
	return false
}

func (account *SavingAccount) GetBalance() float64 {
	return account.balance
}
