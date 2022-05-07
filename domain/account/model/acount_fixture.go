package model

import (
	"github.com/google/uuid"
	customer "go-ddd/domain/customer/model"
)

func CreateNewAccountMock() *Account {
	return CreateNewAccount(uuid.New(), 1010, customer.CreateNewCustomerMock(), 0)
}
