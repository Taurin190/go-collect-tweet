package collect_tweet

import (
	"log"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"github.com/ChimeraCoder/anaconda"
)

func getTwitterApi(secret TwitterSecret) *anaconda.TwitterApi {
    anaconda.SetConsumerKey(secret.ConsumerKey)
    anaconda.SetConsumerSecret(secret.ConsumerSecret)
    return anaconda.NewTwitterApi(secret.AccessToken, secret.AccessSecret)
}

func GetTweetData(api *anaconda.TwitterApi) ([]anaconda.Tweet, error) {
	v := url.Values{}
    v.Set("count", "30")
    searchResult, err := api.GetSearch("golang", v)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return searchResult.Statuses, nil
}

func Exec() ([]anaconda.Tweet, error) {
	tweetSecret, err := GetTwitterSecret()
	if err != nil {
		return nil, err
	}
    api := getTwitterApi(tweetSecret)
	return GetTweetData(api)
}