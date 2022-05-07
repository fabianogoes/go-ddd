package account

import (
	"github.com/google/uuid"
	"go-ddd/domain/account/model"
	customerModel "go-ddd/domain/customer/model"
	"go-ddd/infrastructure/repository"
	"go-ddd/infrastructure/repository/customer"
)

type DBO struct {
	repository.BaseDBO
	Number     int64 `gorm:"column:number"`
	CustomerID uuid.UUID
	Customer   customer.DBO
	Balance    float64 `gorm:"column:balance"`
}

func (DBO) TableName() string {
	return "accounts"
}

func ToModel(dbo *DBO) *model.Account {
	c := customerModel.CreateNewCustomer(dbo.Customer.ID, dbo.Customer.Name, dbo.Customer.CPF)
	return model.CreateNewAccount(dbo.ID, dbo.Number, c, dbo.Balance)
}
