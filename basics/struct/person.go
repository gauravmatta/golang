package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName       string
	LastName        string
	Dob             time.Time
	Email, Location string
}

func (p Person) PrintName() {
	fmt.Printf("\nName: %s %s\n", p.FirstName, p.LastName)
}

func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

func main() {
	var person Person
	person.FirstName = "John"
	person.LastName = "Doe"
	person.Dob = time.Now()
	person.Email = "john.doe@example.com"
	person.Location = "New Delhi"
	person.PrintName()
	person.PrintDetails()

	p1 := Person{
		FirstName: "Jane",
		LastName:  "Smith",
		Dob:       time.Date(1990, time.April, 4, 0, 0, 0, 0, time.UTC),
		Email:     "jane.smith@example.com",
		Location:  "Mumbai",
	}
	p1.PrintName()
	p1.PrintDetails()

	// Pointer to Person Creation by Reference
	p2 := &Person{
		FirstName: "Lampard",
		LastName:  "Smith",
	}
	p2.PrintName()
	p2.PrintDetails()

	// The new built-in function allocates memory
	// The value returned is a pointer to newly allocated type.
	p3 := new(Person)
	p3.FirstName = "Jakarta"
	p3.LastName = "Smith"
	p3.Location = "Mumbai"
	p3.PrintName()
	p3.PrintDetails()
}
