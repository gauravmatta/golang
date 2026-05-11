package main

import "fmt"

var result float64

func doCalculate(num float64, operator string) float64 {
	switch operator {
	case "+":
		return result + num
	case "-":
		return result - num
	case "*":
		return result * num
	case "/":
		return result / num
	default:
		return 0
	}
}
func main() {
	//var result float64
	result = doCalculate(100, "+")
	fmt.Println(result)
	result = doCalculate(50, "-")
	fmt.Println(result)
	result = doCalculate(20, "/")
	fmt.Println(result)
	result = doCalculate(10, "*")
	fmt.Println(result)
}
