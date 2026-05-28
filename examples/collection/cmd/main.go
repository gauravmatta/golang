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
	store.AddItem("country", "India")
	store.AddItem("language", "Go")
	store.AddItem("framework", "Gin")
	fmt.Println(store.GetAll())
	store.UpdateItem("language", "Golang")
	fmt.Println(store.GetAll())
	fmt.Println(store.GetById("language"))
	store.DeleteItem("framework")
	fmt.Println(store.GetAll())
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
