package main

import (
	"customerManager/domain"
	"customerManager/memstore"
	"customerManager/service"
	"fmt"
)

func main() {
	customerController := service.CustomerService{
		Repository: memstore.NewCustomerRepository(),
	}
	customer1 := domain.Customer{
		ID:    "cust01",
		Name:  "Rahul",
		Email: "rahul@gmail.com",
	}
	err := customerController.Create(customer1)
	if err != nil {
		fmt.Println("Error creating customer:", err)
	}
	customers, err := customerController.FindAll()
	if err != nil {
		fmt.Println("Error getting customers:", err)
	} else {
		fmt.Println("Found customers:", customers)
	}
	customer2 := domain.Customer{
		ID:    "cust02",
		Name:  "Mohit",
		Email: "mohit@gmail.com",
	}
	err = customerController.Create(customer2)
	if err != nil {
		fmt.Println("Error creating customer:", err)
	}
	customers, err = customerController.FindAll()
	if err != nil {
		fmt.Println("Error getting customers:", err)
	} else {
		fmt.Println("Found customers:", customers)
	}
	customer3 := domain.Customer{
		ID:    "cust02",
		Name:  "Manu",
		Email: "manu@gmail.com",
	}
	err = customerController.Update("cust02", customer3)
	if err != nil {
		fmt.Println("Error updating customer:", err)
	}
	customers, err = customerController.FindAll()
	if err != nil {
		fmt.Println("Error getting customers:", err)
	} else {
		fmt.Println("Found customers:", customers)
	}
	customer, erf := customerController.FindById("cust01")
	if erf != nil {
		fmt.Println("Error finding customer:", erf)
	} else {
		fmt.Printf("Customer with id 'cust01' was found %s\n", customer)
	}
	err = customerController.Delete("cust02")
	if err != nil {
		fmt.Println("Error deleting customer:", err)
	} else {
		fmt.Println("Customer deleted successfully")
	}
	customers, err = customerController.FindAll()
	if err != nil {
		fmt.Println("Error getting customers:", err)
	} else {
		fmt.Println("Found customers:", customers)
	}
}
