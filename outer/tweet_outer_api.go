package outer

import (
	"../config"
	"../entity"
)

type TweetOuterAPI interface {
	GetTweets() ([]entity.Tweet, error)
}

type tweetOuterAPI struct {
	Config *config.Config
}

func GetTweetOuterAPI(c *config.Config) TweetOuterAPI {
	return &tweetOuterAPI{
		Config: c,
	}
}

func (outerAPI *tweetOuterAPI) GetTweets() ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	return tweets, nil
}
