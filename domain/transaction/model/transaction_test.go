package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction(t *testing.T) {
	t.Run("create new transaction deposit", func(t *testing.T) {
		deposit := CreateNewTransactionDepositMock()
		assert.NotNil(t, deposit)
		assert.NotNil(t, deposit.ID)
		assert.Equalf(t, DEPOSIT.Name(), deposit.TransactionType.Name(), "should be a deposit type")
	})

	t.Run("create new transaction withdraw", func(t *testing.T) {
		deposit := CreateNewTransactionWithdrawMock()
		assert.NotNil(t, deposit)
		assert.NotNil(t, deposit.ID)
		assert.Equalf(t, WITHDRAW.Name(), deposit.TransactionType.Name(), "should be a withdraw type")
	})

	t.Run("create new transaction payment", func(t *testing.T) {
		deposit := CreateNewTransactionPaymentMock()
		assert.NotNil(t, deposit)
		assert.NotNil(t, deposit.ID)
		assert.Equalf(t, PAYMENT.Name(), deposit.TransactionType.Name(), "should be a payment type")
	})

	t.Run("create new transaction transfer", func(t *testing.T) {
		deposit := CreateNewTransactionTransferMock()
		assert.NotNil(t, deposit)
		assert.NotNil(t, deposit.ID)
		assert.Equalf(t, TRANSFER.Name(), deposit.TransactionType.Name(), "should be a deposit type")
	})
}
