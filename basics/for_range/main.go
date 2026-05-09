package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	for i := range 10 {
		fmt.Printf("Index: %d\n", i)
	}
}
