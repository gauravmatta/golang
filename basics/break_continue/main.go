package main

import "fmt"

func main() {
	breakExample()
	continueExapmle()
}

func continueExapmle() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Continuing the loop")
			continue
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting the loop")
}

func breakExample() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking out of the loop")
			break
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exited the Program")
}
