package main

import (
	"fmt"
	"os"
)

func main() {
	ok := CheckFile("input.csv")
	if ok {
		fmt.Println("File correctly read")
	} else {
		fmt.Println("Failed to read file")
	}
}

func CheckFile(path string) bool {
	_, err := os.ReadFile(path)
	return err == nil
}
