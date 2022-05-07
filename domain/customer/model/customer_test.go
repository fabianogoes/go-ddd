package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomer(t *testing.T) {
	t.Run("create new customer", func(t *testing.T) {
		c := CreateNewCustomerMock()
		assert.NotNil(t, c)
		assert.NotNil(t, c.ID)
	})
}
