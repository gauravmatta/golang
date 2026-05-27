package main

import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(5, 3))

	fn := func(principal, interest float64, period int) int {
		var total float64
		total = (principal * interest * float64(period)) / 100
		return int(total)
	}
	result := fn(100, 10, 1)
	fmt.Println(result)
}
