package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-ddd/domain/customer/model"
	customerDBO "go-ddd/infrastructure/repository/customer"
	"go-ddd/infrastructure/repository/database"
	"go-ddd/infrastructure/repository/database/postgres"
	"testing"
)

func TestAccountDataAccess(t *testing.T) {
	ctx := context.Background()
	db, _ := postgres.NewPostgresDB(database.NewPostgresConfigTest())
	if err := db.Migrator().AutoMigrate(&customerDBO.DBO{}, &DBO{}); err != nil {
		panic(err)
	}

	dataAccess := NewAccountDataAccess(db)
	accountNumber := int64(1010)

	t.Run("create a new account", func(t *testing.T) {
		_, err := dataAccess.Create(ctx, accountNumber, model.CreateNewCustomerMock())
		assert.Nil(t, err)

		if found, err := dataAccess.FindAccountByNumber(ctx, accountNumber); err != nil {
			assert.NotNil(t, found.ID)
			assert.Equal(t, accountNumber, found.Number)
		}
	})

	t.Run("find account by number", func(t *testing.T) {
		found, err := dataAccess.FindAccountByNumber(ctx, accountNumber)
		assert.Nil(t, err)
		assert.NotNil(t, found)
		assert.NotNil(t, found.ID)
		assert.Equal(t, accountNumber, found.Number)
	})

	t.Run("update account", func(t *testing.T) {
		found, _ := dataAccess.FindAccountByNumber(ctx, accountNumber)
		depositAmount := float64(1000)
		found.Deposit(depositAmount)

		err := dataAccess.UpdateAccount(ctx, found)
		assert.Nil(t, err)

		updated, _ := dataAccess.FindAccountByNumber(ctx, accountNumber)
		assert.NotNil(t, updated)
		assert.Equal(t, depositAmount, updated.GetBalance())
	})

	_ = db.Migrator().DropTable(&DBO{})
	_ = db.Migrator().DropTable(&customerDBO.DBO{})
}
