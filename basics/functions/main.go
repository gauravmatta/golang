package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func main() {
	x, y := 20, 10
	fmt.Println(x, y)
	fmt.Println(Add(x, y))
	fmt.Println(Subtract(x, y))
}
