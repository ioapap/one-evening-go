package server

import "sync"

type tweetsRepository interface {
	AddTweet(t Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets     []Tweet
	tweetsLock sync.RWMutex
}

func (t *TweetsMemoryRepository) AddTweet(tw Tweet) (int, error) {
	t.tweetsLock.Lock()
	defer t.tweetsLock.Unlock()
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}

func (t *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	t.tweetsLock.RLock()
	defer t.tweetsLock.RUnlock()
	return t.tweets, nil
}
