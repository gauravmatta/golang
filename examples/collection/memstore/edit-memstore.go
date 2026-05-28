package memstore

import (
	"errors"
	"fmt"
)

var data map[string]string

func init() {
	data = make(map[string]string)
}

func AddItem(k, v string) error {
	if _, ok := data[k]; ok {
		return errors.New("duplicate key")
	}
	data[k] = v
	fmt.Println("item added : " + k + "==>" + v)
	return nil
}

func UpdateItem(k, v string) error {
	if _, ok := data[k]; !ok {
		return errors.New("key not found")
	}
	data[k] = v
	return nil
}

func DeleteItem(k string) error {
	if _, ok := data[k]; ok {
		delete(data, k)
		return nil
	}
	return errors.New("key not found")
}
