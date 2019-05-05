package collect_tweet

import (
	"log"
	"time"
	"encoding/json"
	"io/ioutil"
	"gopkg.in/mgo.v2"
)

func GetTwitterSecret(path string) (TwitterSecret, error) {
	var twitterSecret TwitterSecret
	bytes, err := ioutil.ReadFile(path)
    if err != nil {
		return twitterSecret, err
    }
    err2 := json.Unmarshal(bytes, &twitterSecret)
	return twitterSecret, err2
}

func getMongoConf(path string) (MongoConf, error) {
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

func GetMongoInfo(config_path string) (*mgo.DialInfo, error) {
	conf, err := getMongoConf(config_path)
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