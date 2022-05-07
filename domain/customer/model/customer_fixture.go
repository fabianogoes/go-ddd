package model

import "github.com/google/uuid"

func CreateNewCustomerMock() *Customer {
	return CreateNewCustomer(uuid.New(), "Customer test", "123")
}
