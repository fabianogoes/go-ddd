package transaction

import "go-ddd/domain/transaction/model"

type DomainService interface {
	Register(transaction *model.Transaction) (*model.Transaction, error)
}
