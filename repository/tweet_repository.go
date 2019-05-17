package repository

import (
	"time"

	"../config"
	"../entity"
	"gopkg.in/mgo.v2"
)

type TweetRepository interface {
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

func (repository *tweetRepository) Insert(tweets []entity.Tweet) (int, error) {
	count := 0
	info := repository.getDialInfo()
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return count, err
	}
	defer session.Close()
	db := session.DB(repository.Config.MongoDB.Database)

	for _, tweet := range tweets {
		err = db.C(repository.Config.MongoDB.Collection).Insert(tweet)
		if err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

func (repository *tweetRepository) getDialInfo() *mgo.DialInfo {
	info := &mgo.DialInfo{
		Addrs:    []string{repository.Config.MongoDB.Hostname},
		Timeout:  60 * time.Second,
		Database: repository.Config.MongoDB.Database,
		Username: repository.Config.MongoDB.Username,
		Password: repository.Config.MongoDB.Password,
	}
	return info
}
