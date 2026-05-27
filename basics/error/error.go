package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errEmptyFname = errors.New("first name is empty")
	errEmptyLname = errors.New("last name is empty")
	errAgeMinor   = errors.New("age is less than 18")
)

func validate(fname string, lname string, age int) error {
	if len(strings.TrimSpace(fname)) == 0 {
		return errEmptyFname
	}
	if len(strings.TrimSpace(lname)) == 0 {
		return errEmptyLname
	}
	if age < 18 {
		return errAgeMinor
	}
	return nil
}

func getDetails(fname string, lname string, age int) (string, error) {
	if err := validate(fname, lname, age); err != nil {
		return "", err
	}
	return fmt.Sprintf("First Name: %s, Last Name: %s, Age:%d", fname, lname, age), nil
}

func getDetailsWrapErr(fname string, lname string, age int) (string, error) {
	if err := validate(fname, lname, age); err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}
	return fmt.Sprintf("First Name: %s, Last Name: %s, Age:%d", fname, lname, age), nil
}

func main() {
	if str, err := getDetails("Gaurav", "Matta", 40); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(str)
	}
	if str, err := getDetails("Raghavi", "Matta", 7); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(str)
	}
	if _, err := getDetailsWrapErr("Ayra", "Marwah", 3); err != nil {
		fmt.Println(err)
		// Is reports whether any error in err's chain matches target.
		if errors.Is(err, errAgeMinor) {
			fmt.Println("Age is less than 18")
		}
		if err := errors.Unwrap(err); err != nil {
			fmt.Println("UnWrap Error:", err)
		}
	}
}
