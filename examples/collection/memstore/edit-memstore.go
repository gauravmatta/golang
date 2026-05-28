package memstore

import "fmt"

var data map[string]string

func init() {
	data = make(map[string]string)
}

func AddItem(k, v string) {
	if _, ok := data[k]; ok {
		fmt.Println("duplicate key")
		return
	}
	data[k] = v
	fmt.Println("item added : " + k + "==>" + v)
}

func UpdateItem(k, v string) {
	if _, ok := data[k]; !ok {
		fmt.Println("key not found")
		return
	}
	data[k] = v
}

func DeleteItem(k string) {
	delete(data, k)
}
