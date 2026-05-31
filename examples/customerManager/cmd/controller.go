package main

import (
	"customerManager/domain"
	"fmt"
)

type CustomerController struct {
	repository domain.CustomerRepository
}

func (cc CustomerController) Create(customer domain.Customer) error {
	err := cc.repository.Create(customer)
	if err != nil {
		return err
	}
	fmt.Printf("Customer %s has been created\n", customer)
	return nil
}

func (cc CustomerController) Update(id string, customer domain.Customer) error {
	err := cc.repository.Update(id, customer)
	if err != nil {
		return err
	}
	fmt.Printf("Customer %s has been updated\n", customer)
	return nil
}

func (cc CustomerController) Delete(id string) error {
	err := cc.repository.Delete(id)
	if err != nil {
		return err
	}
	fmt.Printf("Customer %s has been deleted\n", id)
	return nil
}

func (cc CustomerController) FindById(id string) (domain.Customer, error) {
	customer, err := cc.repository.FindById(id)
	if err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

func (cc CustomerController) FindAll() ([]domain.Customer, error) {
	customers, err := cc.repository.FindAll()
	if err != nil {
		return []domain.Customer{}, err
	}
	return customers, nil
}
