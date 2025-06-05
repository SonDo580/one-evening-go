package server

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type tweetRepository interface {
	AddTweet(t tweet) (int, error)
	Tweets() ([]tweet, error)
}

type TweetMemoryRepository struct {
	tweets []tweet
}

func (r *TweetMemoryRepository) AddTweet(t tweet) (int, error) {
	r.tweets = append(r.tweets, t)
	return len(r.tweets), nil
}

func (r *TweetMemoryRepository) Tweets() ([]tweet, error) {
	return r.tweets, nil
}
