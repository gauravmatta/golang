package domain

import "errors"

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerRepository interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	FindById(string) (Customer, error)
	FindAll() ([]Customer, error)
}

var (
	ErrNotFound                = errors.New("no records found")
	ErrCustomerExists          = errors.New("customer exists")
	ErrCustomerNotExists error = errors.New("customer doesn't exist")
)
