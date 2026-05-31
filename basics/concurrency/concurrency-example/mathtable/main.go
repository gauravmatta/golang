package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	fmt.Println("Mathtable: Statring routines")
	go addTable(wg)
	go multiTable(wg)
	wg.Wait()
	fmt.Println("Terminating program")
}

func addTable(wg *sync.WaitGroup) {
	defer wg.Done()
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

func multiTable(wg *sync.WaitGroup) {
	defer wg.Done()
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
