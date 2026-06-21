package memstore

import (
	"errors"
	"testing"

	"customerManager/domain"

	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository_CreateValidCustomer(t *testing.T) {
	var wantError error
	customer := domain.Customer{
		ID:    "1",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	wantError = nil
	repository, _ := NewInmemoryRepository()
	err := repository.Create(customer)
	if !errors.Is(err, wantError) {
		t.Errorf("want %v, got %v", wantError, err)
	}
	assert.Equal(t, wantError, err, "failed to create Customer")
}

func TestCustomerRepository_CreateInvalidCustomer(t *testing.T) {
	var wantError error
	customer := domain.Customer{
		ID:    "1",
		Name:  "John Doe",
		Email: "",
	}
	wantError = nil
	repository, _ := NewInmemoryRepository()
	err := repository.Create(customer)
	if !errors.Is(err, wantError) {
		t.Errorf("want %v, got %v", wantError, err)
	}
	assert.Equal(t, wantError, err, "failed to create Customer")
}

func TestCustomerRepository_CreateDuplicateCustomer(t *testing.T) {
	var wantError error
	customer := domain.Customer{
		ID:    "1",
		Name:  "John Doe",
		Email: "john@w.com",
	}
	wantError = nil
	repository, _ := NewInmemoryRepository()
	err := repository.Create(customer)
	customer = domain.Customer{
		ID:    "1",
		Name:  "John Doe",
		Email: "john@w.com",
	}
	err = repository.Create(customer)
	wantError = errors.New("customer with id 1 already exists")
	if !errors.Is(err, wantError) {
		t.Errorf("want %v, got %v", wantError, err)
	}
	assert.Equal(t, wantError, err, "failed to create Customer")
}
