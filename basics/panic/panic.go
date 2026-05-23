package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := ReadFile("test.txt")
	fmt.Println(f)
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	return io.ReadAll(f)
}
