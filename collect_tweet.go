package collect_tweet

import (
	"log"
	"fmt"
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

func LoadConfig(path string) TwitterSecret{
	var tweetConfig TwitterSecret
	bytes, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
    if err := json.Unmarshal(bytes, &tweetConfig); err != nil {
        log.Fatal(err)
    }
	return tweetConfig
}

func getTwitterApi(secret TwitterSecret) *anaconda.TwitterApi {
    anaconda.SetConsumerKey(secret.ConsumerKey)
    anaconda.SetConsumerSecret(secret.ConsumerSecret)
    return anaconda.NewTwitterApi(secret.AccessToken, secret.AccessSecret)
}

func GetTweetData(api *anaconda.TwitterApi) string {
	v := url.Values{}
    v.Set("count", "30")
    searchResult, err := api.GetSearch("golang", v)
	if err != nil {
		log.Fatal(err)
	}
    for _, tweet := range searchResult.Statuses {
        fmt.Println(tweet.Text)
    }
	return ""
}

func Exec() string {
	tweetConfig := LoadConfig("./secret/twitter.conf")
    api := getTwitterApi(tweetConfig)
	return GetTweetData(api)
}