package memstore

import (
	"errors"
	"fmt"
)

var data map[string]string

func init() {
	data = make(map[string]string)
}

func Put(key, value string) error {
	if _, ok := data[key]; ok {
		return fmt.Errorf("key %s already exists", key)
	}
	data[key] = value
	return nil
}

func Get(key string) (string, error) {
	if val, ok := data[key]; ok {
		return val, nil
	}
	return "", errors.New("key not found")
}
