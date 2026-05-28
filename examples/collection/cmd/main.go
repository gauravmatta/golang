package main

import (
	"fmt"

	store "collection/memstore"
)

var data map[string]string

func main() {
	addItem("name", "Gaurav")
	addItem("age", "39")
	addItem("city", "New Delhi")
	fmt.Println(getAll())
	updateItem("name", "Gaurav Matta")
	fmt.Println(getById("name"))
	deleteItem("age")
	fmt.Println(getAll())

	// Using External store library
	fmt.Println("Using External store library")
	if err := store.AddItem("country", "India"); err != nil {
		fmt.Println(err)
	}
	if err := store.AddItem("language", "Go"); err != nil {
		fmt.Println(err)
	}
	if err := store.AddItem("framework", "Gin"); err != nil {
		fmt.Println(err)
	}
	if sl, err := store.GetAll(); err == nil {
		fmt.Println(sl)
	} else {
		fmt.Println(err)
	}
	if err := store.UpdateItem("language", "Golang"); err != nil {
		fmt.Println(err)
	}
	if sl, err := store.GetAll(); err == nil {
		fmt.Println(sl)
	} else {
		fmt.Println(err)
	}
	if val, err := store.GetById("language"); err == nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}
	if err := store.DeleteItem("framework"); err != nil {
		fmt.Println(err)
	}
	if sl, err := store.GetAll(); err == nil {
		fmt.Println(sl)
	} else {
		fmt.Println(err)
	}
}

func init() {
	data = make(map[string]string) // Initialise map with make
}

func addItem(k, v string) {
	if _, ok := data[k]; ok {
		fmt.Println("duplicate key")
		return
	}
	data[k] = v
	fmt.Println("item added")
}

func updateItem(k, v string) {
	if _, ok := data[k]; !ok {
		fmt.Println("key not found")
		return
	}
	data[k] = v
}

func getById(k string) string {
	if _, ok := data[k]; ok {
		fmt.Println("item found")
		return data[k]
	}
	fmt.Println("item not found")
	return ""
}

func getAll() map[string]string {
	return data
}

func deleteItem(k string) {
	delete(data, k)
}
