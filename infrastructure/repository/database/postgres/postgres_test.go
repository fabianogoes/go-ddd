package postgres

import (
	"github.com/stretchr/testify/assert"
	"go-ddd/infrastructure/repository/database"
	"testing"
)

func TestPostgresConnection(t *testing.T) {
	t.Run("connection postgres db", func(t *testing.T) {
		db, err := NewPostgresDB(database.NewPostgresConfigTest())
		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
}
