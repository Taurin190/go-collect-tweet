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

func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func loadConfig() TwitterSecret{
	
}

func getTwitterApi() *anaconda.TwitterApi {
    anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
    anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
    return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func GetTweetData() string {
	loadEnv()
    api := getTwitterApi()
	v := url.Values{}
    v.Set("count", "30")

    searchResult, _ := api.GetSearch("golang", v)
    for _, tweet := range searchResult.Statuses {
        fmt.Println(tweet.Text)
    }
}