package customer

import (
	"testing"
)

func TestCustomer(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		c := &Customer{}
		if c == nil {
			t.Errorf("should be customer not nil %v", c)
		}
	})
}
