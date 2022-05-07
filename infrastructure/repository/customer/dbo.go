package customer

import (
	"go-ddd/domain/customer/model"
	"go-ddd/infrastructure/repository"
)

type DBO struct {
	repository.BaseDBO
	Name string `gorm:"column:name"`
	CPF  string `gorm:"column:cpf"`
}

func (DBO) TableName() string {
	return "customers"
}

func (dbo *DBO) toModel() *model.Customer {
	return &model.Customer{ID: dbo.ID, Name: dbo.Name, CPF: dbo.CPF}
}

func (dbo *DBO) ToDBO(model model.Customer) {
	dbo.ID = model.ID
	dbo.Name = model.Name
	dbo.CPF = model.CPF
}
