package main

import (
	"fmt"

	store "store/services/memstore"
)

func main() {
	if err := store.Put("one", "first"); err != nil {
		fmt.Println("Error putting value:", err)
		return
	}
	if val, err := store.Get("one"); err != nil {
		fmt.Println("Error getting value:", err)
		return
	} else {
		fmt.Println("Retrieved value:", val)
	}
}
