package collect_tweet

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/*
Config structure includes MongoDB and Twitter
*/
type Config struct {
	MongoDB
	Twitter
}

/*
MongoDB Config secret
*/
type MongoDB struct {
	Hostname   string `JSON:"hostname"`
	Database   string `JSON:"database"`
	Collection string `JSON:"collection"`
	Username   string `JSON:"username"`
	Password   string `JSON:"password"`
}

/*
Twitter struct is Twitter secret config
*/
type Twitter struct {
	ConsumerKey    string `JSON:"ConsumerKey"`
	ConsumerSecret string `JSON:"ConsumerSecret"`
	AccessToken    string `JSON:"AccessToken"`
	AccessSecret   string `JSON:"AccessSecret"`
}

/*
GetConfig of structure of config type
*/
func GetConfig() *Config {
	twitterConfig, errTwitter := getTwitterConfig("../conf/twitter.conf")
	mongoConfig, errMongo := getMongoConfig("../conf/mongo.conf")
	if errTwitter != nil {
		log.Fatal(errTwitter)
	}
	if errMongo != nil {
		log.Fatal(errMongo)
	}

	return &Config{
		Twitter: twitterConfig,
		MongoDB: mongoConfig,
	}
}

func getTwitterConfig(path string) (Twitter, error) {
	var twitterConfig Twitter
	bytes, errIo := ioutil.ReadFile(path)
	if errIo != nil {
		return twitterConfig, errIo
	}
	errJSON := json.Unmarshal(bytes, &twitterConfig)
	return twitterConfig, errJSON
}

func getMongoConfig(path string) (MongoDB, error) {
	var mongoConfig MongoDB
	bytes, errIo := ioutil.ReadFile(path)
	if errIo != nil {
		return mongoConfig, errIo
	}
	errJSON := json.Unmarshal(bytes, &mongoConfig)
	return mongoConfig, errJSON
}
