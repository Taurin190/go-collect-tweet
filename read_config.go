package collect_tweet

import (
	"log"
	"time"
	"encoding/json"
	"io/ioutil"
	"gopkg.in/mgo.v2"
)

type TwitterSecret struct {
	ConsumerKey string `JSON:"ConsumerKey"`
	ConsumerSecret string `JSON:"ConsumerSecret"`
	AccessToken string `JSON:"AccessToken"`
	AccessSecret string `JSON:"AccessSecret"`
}

type MongoConf struct {
	hostname string `JSON:"hostname"`
	database string `JSON:"database"`
	collection string `JSON:"collection"`
	username string `JSON:"username"`
	password string `JSON:"password"`
}

func GetTwitterSecret() (TwitterSecret, error) {
	path := "./conf/twitter.conf"
	var twitterSecret TwitterSecret
	bytes, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
		return twitterSecret, err
    }
    if err := json.Unmarshal(bytes, &twitterSecret); err != nil {
        log.Fatal(err)
    }
	return twitterSecret, err
}

func getMongoConf() (MongoConf, error) {
	path := "./conf/mongo.conf"
	var conf MongoConf
	bytes, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
		return conf, err
    }
    if err := json.Unmarshal(bytes, &conf); err != nil {
        log.Fatal(err)
    }
	return conf, err
}

func GetMongoInfo() (*mgo.DialInfo, error) {
	conf, err := getMongoConf()
	if err != nil {
		return nil, err
	}
	info := &mgo.DialInfo{
        Addrs:    []string{conf.hostname},
        Timeout:  60 * time.Second,
        Database: conf.database,
        Username: conf.username,
        Password: conf.password,
    }
	return info, nil
}