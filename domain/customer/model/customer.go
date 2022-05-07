package model

import "github.com/google/uuid"

type Customer struct {
	ID   uuid.UUID
	Name string
	CPF  string
}

func CreateNewCustomer(id uuid.UUID, name string, cpf string) *Customer {
	return &Customer{id, name, cpf}
}
