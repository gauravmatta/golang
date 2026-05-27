package main

import "fmt"

func main() {
	x, y := "Gaurav", "Matta"
	fmt.Println("Before Swap", x, y)
	x, y = Swap(x, y)
	fmt.Println("After Swap", x, y)
	xn, yn := 20, 10
	fmt.Println(xn, yn)
	fmt.Println(Add(xn, yn))
	fmt.Println(Subtract(xn, yn))
}

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Swap(x, y string) (string, string) {
	return y, x
}
