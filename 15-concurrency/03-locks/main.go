package main

import (
	"sync"
	"time"
)

type User struct {
	Email string
}

type Storage struct {
	users map[string]User
}

var lock = &sync.Mutex{}

func (s *Storage) AddUser(email string) {
	lock.Lock()
	defer lock.Unlock()

	_, ok := s.users[email]
	if !ok {
		s.users[email] = User{Email: email}
	}
}

func main() {
	storage := &Storage{users: make(map[string]User)}

	emails := []string{
		"alice@example.com",
		"kate@example.com",
		"joe@example.com",
		"rob@example.com",
		"patrick@example.com",
	}

	for _, email := range emails {
		go storage.AddUser(email)
	}

	time.Sleep(time.Millisecond * 100)
}
