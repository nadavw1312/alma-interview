package dal

import (
	"errors"
	"time"

	"github.com/nadavw1312/golang-fiber/models"
	"github.com/nadavw1312/golang-fiber/utils"
)

type TweetDal struct {
	tweets map[string]*models.Tweet
}

func NewTweetDal() *TweetDal {
	return &TweetDal{tweets: make(map[string]*models.Tweet)}
}

func (dal *TweetDal) InsertTweet(tweet models.NewTweet) (string, error) {
	newTweet := &models.Tweet{UserId: tweet.UserId, Content: tweet.Content, CreatedAt: time.Now()}
	id, err := utils.GenerateRandomID()
	if err != nil {
		return "", err
	}
	newTweet.Id = id
	dal.tweets[id] = newTweet
	return id, nil
}

func (dal *TweetDal) GetTweetById(id string) (*models.Tweet, error) {
	tweet, ok := dal.tweets[id]
	if ok {
		return tweet, nil
	}
	return nil, errors.New("User not found")
}

func (dal *TweetDal) GetTweetsByUserId(userId string) ([]*models.Tweet, error) {
	var tweets []*models.Tweet
	for _, tweet := range dal.tweets {
		if tweet.UserId == userId {
			tweets = append(tweets, tweet)
		}
	}

	return tweets, nil
}
