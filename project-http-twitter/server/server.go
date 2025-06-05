package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	TweetRepository tweetRepository
}

type addTweetResponse struct {
	ID int
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	t := Tweet{}
	if err := json.Unmarshal(body, &t); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", t.Message, t.Location)

	id, err := s.TweetRepository.AddTweet(t)
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

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.TweetRepository.Tweets()
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
