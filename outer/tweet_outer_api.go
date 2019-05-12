package outer

import (
	"../config"
	"../entity"
)

type TweetOuterAPI interface {
}

type tweetOuterAPI struct {
	Config *config.Config
}

func GetTweetOuterAPI(c *config.Config) TweetOuterAPI {
	return &tweetOuterAPI{
		Config: c,
	}
}

func (outerAPI *tweetOuterAPI) GetTweet() (entity.Tweet, error) {
	var tweet entity.Tweet
	return tweet, nil
}
