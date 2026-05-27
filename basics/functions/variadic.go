package main

import "fmt"

func main() {
	total := Sum(1, 2, 3, 4)
	fmt.Println(total)
	total = Sum(5, 7, 8)
	fmt.Println(total)
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	total = Sum(nums...)
	fmt.Println(total)
	total = Sum()
	fmt.Println(total)
}

func Sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
