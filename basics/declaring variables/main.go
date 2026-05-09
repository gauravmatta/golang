package main

import "fmt"

var Title string = "Person Details"

var Country string = "India"

func main() {
	var fname, lname string = "Gaurav", "Matta"
	firstName, lastName := "Raghavi", "Matta" // Short Assignment
	var age int
	var emptybool bool
	var emptyint int
	var emptystring string
	age = 40
	num := 109.9
	ok := true
	const PI = 3.14
	fmt.Println(Title)
	fmt.Println("First Name: ", fname, ", Last Name: ", lastName)
	fmt.Println("Last Name: ", lname, ", First Name: ", firstName)
	fmt.Println("Age: ", age)
	fmt.Println("Country: ", Country)
	fmt.Println(ok)
	fmt.Println(num)
	fmt.Println(emptyint, emptystring, emptybool)
	fmt.Println(PI)
}
