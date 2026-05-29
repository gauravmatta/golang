package main

import (
	"fmt"
	"time"
)

type Worker struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

// PrintName A person method with pointer receiver
func (p *Worker) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

// PrintDetails A person method with pointer receiver
func (p *Worker) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

// ChangeLocation A person method with pointer receiver
func (p *Worker) ChangeLocation(newLocation string) {
	p.Location = newLocation
}

func main() {
	p := &Worker{
		"Gaurav",
		"Matta",
		time.Date(1986, time.April, 4, 0, 0, 0, 0, time.UTC),
		"gm@email.com",
		"Delhi",
	}
	p.ChangeLocation("Santa Clara") // pass by reference
	p.PrintName()
	p.PrintDetails()

	p1 := new(Worker)
	p1.PrintDetails()
}
