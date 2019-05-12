package repository

import (
	"../config"
	"../entity"
)

type TweetRepository interface {
	SetupRepository() error
	Insert(tweets []entity.Tweet) (int, error)
}

type tweetRepository struct {
	Config *config.Config
}

func GetTweetRepository(c *config.Config) TweetRepository {
	return &tweetRepository{
		Config: c,
	}
}

func (repository *tweetRepository) SetupRepository() error {
	return nil
}

func (repository *tweetRepository) Insert(tweets []entity.Tweet) (int, error) {
	return 0, nil
}
