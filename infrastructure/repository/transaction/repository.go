package transaction

import "go-ddd/domain/transaction/model"

type DataAccess struct {
}

func NewTransactionDataAccess() *DataAccess {
	return &DataAccess{}
}

func (tda DataAccess) Register(transaction *model.Transaction) (*model.Transaction, error) {
	// TODO implement
	return model.CreateNewTransactionDepositMock(), nil
}
