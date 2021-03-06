package sqlite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqliteConnection(t *testing.T) {
	t.Run("connection sqlite db", func(t *testing.T) {
		db, err := NewSqliteDB()
		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
}
