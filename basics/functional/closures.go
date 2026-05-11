package main

import "fmt"

func SplitValues(f func(sum int) (int, int)) {
	x, y := f(50)
	fmt.Println(x, y)
}

func main() {
	// A closure is a function value that references variables from outside its body.
	//The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	a, b := 5, 8

	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x
		return x, y
	}

	SplitValues(fn)

	SplitValues(func(sum int) (int, int) {
		return sum / 2, sum / 2
	})
}
