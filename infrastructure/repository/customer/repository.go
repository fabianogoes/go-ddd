package customer

import (
	"context"
	"github.com/google/uuid"
	"go-ddd/domain/customer/model"
	"gorm.io/gorm"
)

type DataAccess struct {
	gormDB *gorm.DB
}

func NewCustomerDataAccess(gormDB *gorm.DB) *DataAccess {
	return &DataAccess{gormDB: gormDB}
}

func (da DataAccess) Create(ctx context.Context, name string, cpf string) (*model.Customer, error) {
	tx := da.gormDB.WithContext(ctx).Begin()

	dbo := &DBO{}
	dbo.ID = uuid.New()
	dbo.Name = name
	dbo.CPF = cpf
	if err := tx.Create(dbo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return dbo.toModel(), tx.Commit().Error
}

func (da DataAccess) FindById(ctx context.Context, id uuid.UUID) (*model.Customer, error) {
	dbo := &DBO{}
	err := da.gormDB.WithContext(ctx).First(dbo, id).Error
	if err != nil {
		return nil, err
	}

	return dbo.toModel(), nil
}

func (da DataAccess) FindByCPF(ctx context.Context, cpf string) (*model.Customer, error) {
	dbo := &DBO{}
	err := da.gormDB.WithContext(ctx).Where("cpf = ?", cpf).First(dbo).Error
	if err != nil {
		return nil, err
	}

	return dbo.toModel(), nil
}
