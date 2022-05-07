package model

import (
	"github.com/google/uuid"
	"go-ddd/domain/customer/model"
)

type Account struct {
	ID       uuid.UUID
	Number   int64
	Customer *model.Customer
	balance  float64
}

func CreateNewAccount(id uuid.UUID, number int64, customer *model.Customer, currentBalance float64) *Account {
	return &Account{id, number, customer, currentBalance}
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) *Account {
	a.balance += amount
	return a
}
