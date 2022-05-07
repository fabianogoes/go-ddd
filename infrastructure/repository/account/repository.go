package account

import (
	"context"
	"github.com/google/uuid"
	"go-ddd/domain/account/model"
	customer "go-ddd/domain/customer/model"
	customerDBO "go-ddd/infrastructure/repository/customer"
	"gorm.io/gorm"
)

type DataAccess struct {
	gormDB *gorm.DB
}

func NewAccountDataAccess(db *gorm.DB) *DataAccess {
	return &DataAccess{gormDB: db}
}

func (da DataAccess) Create(ctx context.Context, number int64, customer *customer.Customer) (*model.Account, error) {
	tx := da.gormDB.WithContext(ctx).Begin()

	dbo := &DBO{}
	dbo.ID = uuid.New()
	dbo.Number = number
	dbo.Customer = customerDBO.DBO{Name: customer.Name, CPF: customer.CPF}
	if err := tx.Create(dbo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return ToModel(dbo), tx.Commit().Error
}

func (da DataAccess) FindAccountByNumber(ctx context.Context, number int64) (*model.Account, error) {
	dbo := &DBO{}
	err := da.gormDB.WithContext(ctx).Where("number = ?", number).First(dbo).Error
	if err != nil {
		return nil, err
	}

	return ToModel(dbo), nil
}

func (da DataAccess) UpdateAccount(ctx context.Context, account *model.Account) error {
	tx := da.gormDB.WithContext(ctx).Begin()

	dbo := &DBO{}
	dbo.Number = account.Number
	dbo.Customer.ToDBO(*account.Customer)
	dbo.Balance = account.GetBalance()
	if err := tx.Create(dbo).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
