package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount(t *testing.T) {
	t.Run("create new account", func(t *testing.T) {
		a := CreateNewAccountMock()
		assert.NotNil(t, a)
		assert.NotNil(t, a.ID)
		assert.NotNil(t, a.Customer)
		assert.NotNil(t, a.Customer.ID)
	})
}
