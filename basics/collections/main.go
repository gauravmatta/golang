package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[0] = 100
	x[4] = 25
	fmt.Println(x)
	x[1] = 10
	x[2] = 15
	x[3] = 20
	fmt.Println(x)

	// Declare and initialize an array in one line with array literal
	y := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Value of y:", y)
	fmt.Println("Length of y:", len(y))

	// Array literal with ...
	z := [...]int{10, 20, 30, 40, 50}
	fmt.Println("Value of z:", z)
	fmt.Println("Length of z:", len(z))

	//Initialize at specific index with array literal
	langs := [4]string{0: "Go", 3: "Rust"}
	fmt.Println("Value of langs:", langs)
	langs[1] = "Julia"
	langs[2] = "Scala"
	fmt.Println("Value of langs:", langs)
	fmt.Println("Length of langs:", len(langs))
	fmt.Println("\nIterate over arrays\n")

	for i := 0; i < len(langs); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, langs[i])
	}

	fmt.Println("\nIterate over arrays using for range\n")
	for k, v := range langs {
		fmt.Printf("Index: %d, Value: %s\n", k, v)
	}

	fmt.Println("\nIterate over keys using for range\n")
	for k := range langs {
		fmt.Printf("Index: %d, Value: %s\n", k, langs[k])
	}

	fmt.Println("\nIterate over values using for range\n")
	for _, v := range langs {
		fmt.Printf("Value: %s\n", v)
	}
}
