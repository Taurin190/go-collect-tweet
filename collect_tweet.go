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

func loadConfig() TwitterSecret{
	var tweetConfig TwitterSecret
	bytes, err := ioutil.ReadFile("./secret/twitter.conf")
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

func GetTweetData() string {
	tweetConfig := loadConfig()
    api := getTwitterApi(tweetConfig)
	v := url.Values{}
    v.Set("count", "30")

    searchResult, _ := api.GetSearch("golang", v)
    for _, tweet := range searchResult.Statuses {
        fmt.Println(tweet.Text)
    }

	return ""
}