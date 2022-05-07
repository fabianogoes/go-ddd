package model

import (
	"github.com/google/uuid"
	"go-ddd/domain/account/model"
)

type Transaction struct {
	ID              uuid.UUID
	Account         *model.Account
	Amount          float64
	TransactionType TransactionType
}

func CreateNewTransaction(account *model.Account, amount float64, transactionType TransactionType) *Transaction {
	return &Transaction{uuid.New(), account, amount, transactionType}
}
