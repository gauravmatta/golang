package main

import "fmt"

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
}

func init() {
	data = make(map[string]string) // Initialise map with make
}

func addItem(k, v string) {
	data[k] = v
}

func updateItem(k, v string) {
	data[k] = v
}

func getById(k string) string {
	return data[k]
}

func getAll() map[string]string {
	return data
}

func deleteItem(k string) {
	delete(data, k)
}
