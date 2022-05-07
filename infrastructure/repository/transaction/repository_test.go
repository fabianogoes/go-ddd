package transaction

import (
	"github.com/stretchr/testify/assert"
	"go-ddd/domain/transaction/model"
	"testing"
)

func TestTransactionDataAccess(t *testing.T) {
	t.Run("register deposit transaction", func(t *testing.T) {
		dataAccess := NewTransactionDataAccess()

		registered, err := dataAccess.Register(model.CreateNewTransactionDepositMock())
		assert.Nil(t, err)
		assert.NotNil(t, registered)
		assert.NotNil(t, registered.ID)
		assert.Equal(t, model.DEPOSIT.Name(), registered.TransactionType.Name())
	})
}
