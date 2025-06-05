package main

import (
	"errors"
)

var message = ""

func StoreMessage(m string) error {
	if m == "" {
		return errors.New("no message")
	}

	message = m

	return nil
}

func main() {
	MustStoreMessage("Hello!")
}

func MustStoreMessage(msg string) {
	err := StoreMessage(msg)
	if err != nil {
		panic(err)
	}
}
