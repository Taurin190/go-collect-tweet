package outer

import (
	"log"
	"net/url"

	"../config"
	"../entity"
	"github.com/ChimeraCoder/anaconda"
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
	tweets, err := outerAPI.getTweetData()
	return tweets, err
}

func (outerAPI *tweetOuterAPI) getTweetData() ([]entity.Tweet, error) {
	anaconda.SetConsumerKey(outerAPI.Config.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(outerAPI.Config.Twitter.ConsumerSecret)
	api := anaconda.NewTwitterApi(outerAPI.Config.Twitter.AccessToken, outerAPI.Config.Twitter.AccessSecret)
	v := url.Values{}
	v.Set("count", "30")
	searchResult, err := api.GetSearch("golang", v)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return entity.ConvertTweetsStruct(searchResult.Statuses), nil
}
