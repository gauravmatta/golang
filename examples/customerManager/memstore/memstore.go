package memstore

import (
	"fmt"

	"customerManager/domain"
)

type CustomerRepository struct {
	customers map[string]domain.Customer
}

func (m *CustomerRepository) Create(customer domain.Customer) error {
	m.customers[customer.ID] = customer
	return nil
}

func (m *CustomerRepository) Update(id string, customer domain.Customer) error {
	m.customers[id] = customer
	return nil
}

func (m *CustomerRepository) Delete(id string) error {
	delete(m.customers, id)
	return nil
}

func (m *CustomerRepository) FindById(id string) (domain.Customer, error) {
	customer, exists := m.customers[id]
	if !exists {
		return domain.Customer{}, fmt.Errorf("customer not found")
	}
	return customer, nil
}

func (m *CustomerRepository) FindAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	for _, customer := range m.customers {
		customers = append(customers, customer)
	}
	return customers, nil
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{customers: make(map[string]domain.Customer)}
}
