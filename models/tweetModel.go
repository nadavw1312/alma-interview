package models

import "time"

type Tweet struct {
	Id        string    `json:"id"`
	UserId    string    `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewTweet struct {
	UserId    string    `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

// Implement the sort.Interface for your slice
type TweetsByDate []*Tweet

func (a TweetsByDate) Len() int           { return len(a) }
func (a TweetsByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TweetsByDate) Less(i, j int) bool { return a[i].CreatedAt.After(a[j].CreatedAt) }
