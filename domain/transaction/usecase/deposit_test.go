package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-ddd/domain/customer/model"
	"go-ddd/infrastructure/repository"
	"go-ddd/infrastructure/repository/account"
	customerDBO "go-ddd/infrastructure/repository/customer"
	"go-ddd/infrastructure/repository/transaction"
	"testing"
)

func TestTransactionUseCase(t *testing.T) {
	ctx := context.Background()
	db, _ := repository.NewPostgresDB(repository.NewConfigDBTest())
	if err := db.Migrator().AutoMigrate(&customerDBO.DBO{}, &account.DBO{}); err != nil {
		panic(err)
	}

	dataAccess := account.NewAccountDataAccess(db)
	accountNumber := int64(1010)

	t.Run("deposit use case", func(t *testing.T) {
		a, errCreate := dataAccess.Create(ctx, accountNumber, model.CreateNewCustomerMock())
		assert.Nil(t, errCreate)
		assert.NotNil(t, a)

		ads := account.NewAccountDataAccess(db)
		tds := transaction.NewTransactionDataAccess()
		useCase := NewDepositUseCase(ads, tds)
		accountNumber := int64(1010)
		depositAmount := 100.50

		err := useCase.deposit(ctx, accountNumber, depositAmount)

		assert.Nil(t, err)
	})

	_ = db.Migrator().DropTable(&account.DBO{})
	_ = db.Migrator().DropTable(&customerDBO.DBO{})
}
