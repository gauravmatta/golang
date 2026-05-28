package memstore

import "errors"

func GetById(k string) (string, error) {
	if v, ok := data[k]; ok {
		return v, nil
	}
	return "", errors.New("key not found")
}

func GetAll() ([]string, error) {
	if len(data) == 0 {
		return nil, errors.New("data store is empty")
	}
	values := make([]string, 0)
	for k, v := range data {
		values = append(values, k+"=>"+v)
	}
	return values, nil
}
