package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostgres(t *testing.T) {
	t.Run("connection postgres db", func(t *testing.T) {
		db, err := NewPostgresDB(NewConfigDBTest())
		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
}
