package main

import "fmt"

func main() {
	counter := make(chan int)
	nums := make(chan int, 3)
	go func() {
		counter <- 1
		close(counter)
	}()
	go func() {
		nums <- 20
		nums <- 30
		nums <- 40
		nums <- 50
	}()
	nums <- 60
	nums <- 70
	nums <- 80
	v, exists := <-counter
	fmt.Printf("Counter: %d, exists: %t\n", v, exists)
	v2, exists2 := <-nums
	fmt.Printf("Nums: %d, exists: %t\n", v2, exists2)
	val, ok := <-counter
	fmt.Printf("Counter: %d, ok: %t\n", val, ok)
	l1 := len(counter)
	l2 := len(nums)
	fmt.Printf("Length of counter: %d, Length of nums: %d\n", l1, l2)
	fmt.Println(<-nums)
	fmt.Println(<-nums)
	fmt.Println(<-nums)
	fmt.Println(<-nums)
	fmt.Println(<-nums)
	close(nums)
}
