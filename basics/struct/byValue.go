package main

import (
	"fmt"
	"time"
)

type Character struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

// PrintName A character method with value receiver
func (c Character) PrintName() {
	fmt.Printf("\n%s %s\n", c.FirstName, c.LastName)
}

// PrintDetails A character method with value receiver
func (c Character) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", c.Dob.String(), c.Email, c.Location)
}

// ChangeLocation A character method with value receiverz
func (c Character) ChangeLocation(newLocation string) {
	c.Location = newLocation // mutating location
	fmt.Println(c.Location)
}

func main() {
	p := Character{
		"Gaurav",
		"Matta",
		time.Date(1986, time.April, 4, 0, 0, 0, 0, time.UTC),
		"gaurav@email.com",
		"Delhi",
	}
	p.ChangeLocation("Santa Clara")
	p.PrintName()
	p.PrintDetails() // Location won't change
}
