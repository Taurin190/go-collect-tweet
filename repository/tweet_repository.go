package repository

import (
	"../config"
	"../entity"
)

type TweetRepository interface {
}

type tweetRepository struct {
	Config *config.Config
}

func GetTweetRepository(c *config.Config) TweetRepository {
	return &tweetRepository{
		Config: c,
	}
}

func (repository *tweetRepository) Insert(tweet entity.Tweet) error {

	return nil
}
