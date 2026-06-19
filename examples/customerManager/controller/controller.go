package controller

import (
	"customerManager/domain"
	"fmt"
)

type CustomerController struct {
	Repository domain.CustomerRepository
}

func (cc CustomerController) Create(customer domain.Customer) error {
	err := cc.Repository.Create(customer)
	if err != nil {
		return domain.ErrCustomerExists
	}
	fmt.Printf("Customer %s has been created\n", customer)
	return nil
}

func (cc CustomerController) Update(id string, customer domain.Customer) error {
	err := cc.Repository.Update(id, customer)
	if err != nil {
		return domain.ErrCustomerNotExists
	}
	fmt.Printf("Customer %s has been updated\n", customer)
	return nil
}

func (cc CustomerController) Delete(id string) error {
	err := cc.Repository.Delete(id)
	if err != nil {
		return domain.ErrCustomerNotExists
	}
	fmt.Printf("Customer %s has been deleted\n", id)
	return nil
}

func (cc CustomerController) FindById(id string) (domain.Customer, error) {
	customer, err := cc.Repository.FindById(id)
	if err != nil {
		return domain.Customer{}, domain.ErrCustomerNotExists
	}
	return customer, nil
}

func (cc CustomerController) FindAll() ([]domain.Customer, error) {
	customers, err := cc.Repository.FindAll()
	if err != nil {
		return []domain.Customer{}, domain.ErrNotFound
	}
	return customers, nil
}
