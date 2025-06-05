package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	s := server.Server{
		TweetRepository: &server.TweetMemoryRepository{},
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/tweets", s.ListTweets)
	r.With(httprate.LimitByIP(10, time.Minute)).Post("/tweets", s.AddTweet)

	go spamTweets()

	log.Fatal(http.ListenAndServe(":8080", r))
}

func spamTweets() {
	url := "http://localhost:8080/tweets"

	addTweetPayload := server.Tweet{
		Message:  "ass",
		Location: "ass",
	}

	marshaledPayload, err := json.Marshal(addTweetPayload)
	if err != nil {
		panic(err)
	}

	for {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Failed to read response body", err)
		}
		resp.Body.Close()

		fmt.Println(string(body))
	}
}
