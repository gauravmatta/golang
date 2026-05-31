package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Mathtable: Statring routines")
	addTable()
	multiTable()
	fmt.Println("Terminating program")
}

func addTable() {
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Adding ", i)
		for j := 1; j <= 10; j++ {
			fmt.Println("Adding ", i, " + ", j, " = ", i+j)
		}
		fmt.Println("Finished adding ", i)
	}
}

func multiTable() {
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Multiplying ", i)
		for j := 1; j <= 10; j++ {
			fmt.Println("Multiplying ", i, " * ", j, " = ", i*j)
		}
		fmt.Println("Finished multiplying ", i)
	}
}
