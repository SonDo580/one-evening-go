package main

import "fmt"

func Alphabet(length int) []string {
	letters := []string{}
	for i := 0; i < length; i++ {
		letters = append(letters, characterByIndex(i))
	}
	return letters
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
