package collect_tweet

import (
	"log"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	MongoDB
	Twitter
}

type MongoDB struct {
	Hostname string `JSON:"hostname"`
	Database string `JSON:"database"`
	Collection string `JSON:"collection"`
	Username string `JSON:"username"`
	Password string `JSON:"password"`
}

type Twitter struct {
	ConsumerKey string `JSON:"ConsumerKey"`
	ConsumerSecret string `JSON:"ConsumerSecret"`
	AccessToken string `JSON:"AccessToken"`
	AccessSecret string `JSON:"AccessSecret"`
}

func GetConfig() *Config {
	twitterConfig, err_twitter := getTwitterConfig("../conf/twitter.conf")
	mongoConfig, err_mongo := getMongoConfig("../conf/mongo.conf")
	if err_twitter != nil {
		log.Fatal(err_twitter)
	}
	if err_mongo != nil {
		log.Fatal(err_mongo)
	}

	return &Config {
		Twitter: twitterConfig,
		MongoDB: mongoConfig,
	}
}

func getTwitterConfig(path string) (Twitter, error) {
	var twitterConfig Twitter
	bytes, err_io := ioutil.ReadFile(path)
    if err_io != nil {
		return twitterConfig, err_io
    }
    err_json := json.Unmarshal(bytes, &twitterConfig)
	return twitterConfig, err_json
}

func getMongoConfig(path string) (MongoDB, error) {
	var mongoConfig MongoDB
	bytes, err_io := ioutil.ReadFile(path)
    if err_io != nil {
		return mongoConfig, err_io
    }
    err_json := json.Unmarshal(bytes, &mongoConfig)
	return mongoConfig, err_json
}