package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	s := server{
		tweetRepository: &tweetMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type server struct {
	tweetRepository tweetRepository
}

func (s server) tweets(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	if r.Method == http.MethodGet {
		s.listTweets(w, r)
	} else if r.Method == http.MethodPost {
		s.addTweet(w, r)
	}
}

func (s server) addTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	t := tweet{}
	if err := json.Unmarshal(body, &t); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", t.Message, t.Location)

	id, err := s.tweetRepository.AddTweet(t)
	if err != nil {
		log.Println("Failed to add tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(addTweetResponse{ID: id})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}

func (s server) listTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.tweetRepository.Tweets()
	if err != nil {
		log.Println("Failed to retrieve tweets:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(tweetsList{Tweets: tweets})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type addTweetResponse struct {
	ID int
}

type tweetsList struct {
	Tweets []tweet `json:"tweets"`
}

type tweetRepository interface {
	AddTweet(t tweet) (int, error)
	Tweets() ([]tweet, error)
}

type tweetMemoryRepository struct {
	tweets []tweet
}

func (r *tweetMemoryRepository) AddTweet(t tweet) (int, error) {
	r.tweets = append(r.tweets, t)
	return len(r.tweets), nil
}

func (r *tweetMemoryRepository) Tweets() ([]tweet, error) {
	return r.tweets, nil
}
