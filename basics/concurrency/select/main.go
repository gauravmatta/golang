package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type fibvalue struct {
	input, value int
}
type squarevalue struct {
	input, value int
}

var wg sync.WaitGroup

func generateSquares(sqrs chan<- squarevalue) {
	defer wg.Done()
	defer close(sqrs)
	for i := 1; i <= 10; i++ {
		num := rand.Intn(50)
		sqrs <- squarevalue{
			input: num,
			value: num * num,
		}
	}
}

func generateFibonacci(fibs chan<- fibvalue) {
	defer wg.Done()
	defer close(fibs)
	for i := 1; i <= 10; i++ {
		num := float64(rand.Intn(50))
		// Fibonacci using Binet's formula
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, num) - math.Pow(phi, num)) / math.Sqrt(5)
		fibs <- fibvalue{
			input: int(num),
			value: int(result),
		}
	}
}

func printValues(fibs <-chan fibvalue, sqrs <-chan squarevalue) {
	defer wg.Done()
	for fibs != nil || sqrs != nil {
		select {
		case fib, ok := <-fibs:
			if ok {
				fmt.Printf("Fibonacci value of %d is %d\n", fib.input, fib.value)
			} else {
				fibs = nil
				fmt.Println("Fibonacci channel closed")
			}
		case sqr, ok := <-sqrs:
			if ok {
				fmt.Printf("Square value of %d is %d\n", sqr.input, sqr.value)
			} else {
				sqrs = nil
				fmt.Println("Square channel closed")
			}
		case <-time.After(500 * time.Millisecond):
			fmt.Println("No value received in the last 500ms")
		}
	}
}

func main() {
	fibs := make(chan fibvalue)
	sqrs := make(chan squarevalue)
	wg.Add(3)
	go generateFibonacci(fibs)
	go generateSquares(sqrs)
	go printValues(fibs, sqrs)
	wg.Wait()
}
