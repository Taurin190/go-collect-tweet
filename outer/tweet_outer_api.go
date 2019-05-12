package outer

import (
	"../entity"
)

type TweetOuterAPI interface {
}

type tweetOuterAPI struct {
}

func GetTweetOuterAPI() TweetOuterAPI {
	return &tweetOuterAPI{}
}

func (outerAPI *tweetOuterAPI) GetTweet() (entity.Tweet, error) {
	var tweet entity.Tweet
	return tweet, nil
}
