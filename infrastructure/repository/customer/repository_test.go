package customer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-ddd/infrastructure/repository"
	"testing"
)

const (
	name = "test"
	cpf  = "123"
)

func TestCustomerRepository(t *testing.T) {
	ctx := context.Background()
	db, _ := repository.NewPostgresDB(repository.NewPostgresConfigTest())
	_ = db.Migrator().CreateTable(&DBO{})
	dataAccess := NewCustomerDataAccess(db)

	t.Run("create a new customer", func(t *testing.T) {
		found, err := dataAccess.Create(ctx, name, cpf)
		assert.Nil(t, err)
		assert.NotNil(t, found)
		assert.NotNil(t, found.ID)
		assert.Equal(t, cpf, found.CPF)
		assert.Equal(t, name, found.Name)
	})

	t.Run("find by cpf", func(t *testing.T) {
		created, err := dataAccess.Create(ctx, name, cpf)
		assert.Nil(t, err)

		found, err := dataAccess.FindByCPF(ctx, cpf)
		assert.Nil(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, created.CPF, found.CPF)
	})

	t.Run("find by id", func(t *testing.T) {
		created, err := dataAccess.Create(ctx, name, cpf)
		assert.Nil(t, err)

		found, err := dataAccess.FindById(ctx, created.ID)
		assert.Nil(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, created.ID, found.ID)
	})

	_ = db.Migrator().DropTable(&DBO{})
}
