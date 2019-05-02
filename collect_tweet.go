package collect_tweet

import (
	"io/ioutil"
	"github.com/ChimeraCoder/anaconda"
)

type TwitterSecret struct {
	CONSUMER_KEY string `JSON:"ConsumerKey"`
	CONSUMER_SECRET string `JSON:"ConsumerSecret"`
	ACCESS_TOKEN string `JSON:"AccessToken"`
	ACCESS_TOKEN_SECRET string `JSON:"AccessSecret"`
}

func loadConfig() TwitterSecret{
	tweetConfig := []TwitterSecret
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
    anaconda.SetConsumerKey(secret("CONSUMER_KEY"))
    anaconda.SetConsumerSecret(secret("CONSUMER_SECRET"))
    return anaconda.NewTwitterApi(secret("ACCESS_TOKEN"), secret("ACCESS_TOKEN_SECRET"))
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
}