package model

import "go-ddd/domain/account/model"

func CreateNewTransactionDepositMock() *Transaction {
	return CreateNewTransaction(model.CreateNewAccountMock(), 1000, DEPOSIT)
}

func CreateNewTransactionWithdrawMock() *Transaction {
	return CreateNewTransaction(model.CreateNewAccountMock(), 100, WITHDRAW)
}

func CreateNewTransactionPaymentMock() *Transaction {
	return CreateNewTransaction(model.CreateNewAccountMock(), 150, PAYMENT)
}

func CreateNewTransactionTransferMock() *Transaction {
	return CreateNewTransaction(model.CreateNewAccountMock(), 200, TRANSFER)
}
