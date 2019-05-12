package repository

import (
	"../config"
	"../entity"
)

type TweetRepository interface {
}

type tweetRepository struct {
	config config.Config
}

func GetTweetRepository() TweetRepository {
	return &tweetRepository{}
}

func (repository *tweetRepository) Insert(tweet entity.Tweet) error {

	return nil
}
