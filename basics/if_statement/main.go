package main

import "fmt"

func main() {
	tenCheck()
	evenCheck(8)
	evenCheck(5)
}

func evenCheck(input int) {
	if input%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}

func tenCheck() {
	x := 5
	if x < 10 {
		fmt.Println("x is less than 10")
	}
}
