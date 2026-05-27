package main

import "fmt"

func main() {
	var sl []string // nil slice
	x := make([]int, 3, 10)
	x[0] = 10
	x[1] = 20
	x[2] = 30
	//x[3] = 40
	println(x, "x")

	y := make([]int, 3)
	y[0] = 10
	y[1] = 20
	y[2] = 30
	println(y, "y")

	z := []int{10, 20, 30}
	println(z, "z")

	z1 := []int{0: 10, 2: 30}
	println(z1, "z1")

	slice := []int{10, 20, 30}
	slice = append(slice, 50)
	slice = append(slice, []int{30, 40, 50}...)
	println(slice, "slice")

	sliced := slice[1:3]
	println(sliced, "sliced")

	sl = append(sl, "Go", "Rust", "Python")
	prints(sl, "sl")
}

func println(s []int, label string) {
	fmt.Println()
	fmt.Println("Slice:", label)
	fmt.Println("Array", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
}

func prints(s []string, label string) {
	fmt.Println()
	fmt.Println("Slice:", label)
	fmt.Println("Array", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
}
