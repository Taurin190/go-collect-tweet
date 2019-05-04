package collect_tweet

import (
	"log"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"github.com/ChimeraCoder/anaconda"
)

type TwitterSecret struct {
	ConsumerKey string `JSON:"ConsumerKey"`
	ConsumerSecret string `JSON:"ConsumerSecret"`
	AccessToken string `JSON:"AccessToken"`
	AccessSecret string `JSON:"AccessSecret"`
}

func LoadConfig(path string) (TwitterSecret, error) {
	var tweetConfig TwitterSecret
	bytes, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
		return tweetConfig, err
    }
    if err := json.Unmarshal(bytes, &tweetConfig); err != nil {
        log.Fatal(err)
    }
	return tweetConfig, err
}

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
	tweetConfig, err := LoadConfig("./secret/twitter.conf")
	if err != nil {
		return nil, err
	}
    api := getTwitterApi(tweetConfig)
	return GetTweetData(api)
}