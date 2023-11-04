package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errZeroMoney = errors.New("[ERR] MONEY ZERO")

// NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0,
	}
	return &account
}

// Deposit +amount
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// GetBalance get balance
func (a Account) GetBalance() int {
	return a.balance
}

// Withdraw -amount
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errZeroMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner change owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) GetOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprintf("This is %s's account\n", a.owner)
}
