package main

import "fmt"

func main() {
	getType(25)
	getType(true)
	getType("Gopher")
}

func getType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
