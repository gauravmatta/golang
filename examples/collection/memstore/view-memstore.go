package memstore

func GetById(k string) string {
	if _, ok := data[k]; ok {
		return data[k]
	}
	return ""
}

func GetAll() map[string]string {
	return data
}
