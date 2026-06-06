package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type fibvalue1 struct {
	input, value int
}

var wg1 sync.WaitGroup

func randomCounterWithChanOper(out chan<- int) {
	defer wg1.Done()
	var random int
	for x := 0; x < 10; x++ {
		random = rand.Intn(50)
		out <- random
	}
	close(out)
}

func generateFibonacciWithChanOper(out chan<- fibvalue1, in <-chan int) {
	defer wg1.Done()
	var input float64
	for v := range in {
		input = float64(v)
		// Fibonacci using Binet's formula
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, input) - math.Pow(phi, input)) / math.Sqrt(5)
		out <- fibvalue1{
			input: v,
			value: int(result),
		}
	}
	close(out)

	/*
			for {
				v, ok:= <-in
				if !ok {
					fmt.Println("Channel was closed")
					return
				}
				// find out fibonacci value using v
				// Once fibonacci value is generated
				out <- fibvalue{
					input: v,
					value:0, //fibonacci value
				}
		  }
	*/
}

func printFibonacciWithChanOper(in <-chan fibvalue1) {
	defer wg1.Done()
	for v := range in {
		fmt.Printf("Fibonacci value of %d is %d\n", v.input, v.value)
	}
}

func main() {
	// Add 3 into WaitGroup Counter
	wg1.Add(3)
	// Declare Channels
	randoms := make(chan int)
	fibs := make(chan fibvalue1)
	// Launching 3 goroutines
	go randomCounterWithChanOper(randoms)
	go generateFibonacciWithChanOper(fibs, randoms)
	go printFibonacciWithChanOper(fibs)
	// Wait for completing all goroutines
	wg1.Wait()
}
