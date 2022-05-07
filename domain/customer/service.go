package customer

import (
	"context"
	"github.com/google/uuid"
	"go-ddd/domain/customer/model"
)

type DomainService interface {
	Create(ctx context.Context, name string, cpf string) (*model.Customer, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.Customer, error)
	FindByCPF(ctx context.Context, cpf string) (*model.Customer, error)
}
