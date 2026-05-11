package main

import "fmt"

func main() {

	x, y := "Gaurav", "Matta"
	fmt.Println("Before Swap", x, y)
	x, y = Swap(x, y)
	fmt.Println("After Swap", x, y)
}

func Swap(x, y string) (string, string) {
	return y, x
}
