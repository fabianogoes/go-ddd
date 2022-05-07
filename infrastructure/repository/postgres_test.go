package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresConnection(t *testing.T) {
	t.Run("connection postgres db", func(t *testing.T) {
		db, err := NewPostgresDB(NewPostgresConfigTest())
		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
}
