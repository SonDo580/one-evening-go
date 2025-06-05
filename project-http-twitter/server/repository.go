package server

import (
	"sync"
)

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type tweetRepository interface {
	AddTweet(t Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetMemoryRepository struct {
	tweets []Tweet
	lock   sync.RWMutex
}

func (r *TweetMemoryRepository) AddTweet(t Tweet) (int, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.tweets = append(r.tweets, t)
	return len(r.tweets), nil
}

func (r *TweetMemoryRepository) Tweets() ([]Tweet, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.tweets, nil
}
