package collect_tweet

import (
	"github.com/ChimeraCoder/anaconda"
)

type TwitterSecret struct {
	ConsumerKey string `JSON:"ConsumerKey"`
	ConsumerSecret string `JSON:"ConsumerSecret"`
	AccessToken string `JSON:"AccessToken"`
	AccessSecret string `JSON:"AccessSecret"`
}

func loadConfig() TwitterSecret{
	
}

func getTwitterApi(secret TwitterSecret) *anaconda.TwitterApi {
    anaconda.SetConsumerKey(secret("CONSUMER_KEY"))
    anaconda.SetConsumerSecret(secret("CONSUMER_SECRET"))
    return anaconda.NewTwitterApi(secret("ACCESS_TOKEN"), secret("ACCESS_TOKEN_SECRET"))
}

func GetTweetData() string {
	tweetConfig := loadConfig()
    api := getTwitterApi()
	v := url.Values{}
    v.Set("count", "30")

    searchResult, _ := api.GetSearch("golang", v)
    for _, tweet := range searchResult.Statuses {
        fmt.Println(tweet.Text)
    }
}