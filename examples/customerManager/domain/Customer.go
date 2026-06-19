package domain

import "errors"

type Customer struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type CustomerRepository interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	FindById(string) (Customer, error)
	FindAll() ([]Customer, error)
}

var (
	ErrNotFound          = errors.New("no records found")
	ErrCustomerExists    = errors.New("customer exists")
	ErrCustomerNotExists = errors.New("customer doesn't exist")
)
