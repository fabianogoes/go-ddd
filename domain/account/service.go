package account

import (
	"context"
	"go-ddd/domain/account/model"
	customer "go-ddd/domain/customer/model"
)

type DomainService interface {
	Create(ctx context.Context, number int64, customer *customer.Customer) (*model.Account, error)
	FindAccountByNumber(ctx context.Context, number int64) (*model.Account, error)
	UpdateAccount(ctx context.Context, account *model.Account) error
}
