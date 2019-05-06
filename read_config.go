package collect_tweet

import (
	"time"
	"encoding/json"
	"io/ioutil"
	"gopkg.in/mgo.v2"
)

func GetTwitterSecret(path string) (TwitterSecret, error) {
	var twitterSecret TwitterSecret
	bytes, err_io := ioutil.ReadFile(path)
    if err_io != nil {
		return twitterSecret, err_io
    }
    err_json := json.Unmarshal(bytes, &twitterSecret)
	return twitterSecret, err_json
}

func getMongoConf(path string) (MongoConf, error) {
	var mongoConf MongoConf
	bytes, err_io := ioutil.ReadFile(path)
    if err_io != nil {
		return mongoConf, err_io
    }
    err_json := json.Unmarshal(bytes, &mongoConf)
	return mongoConf, err_json
}

func GetMongoInfo(path string) (*mgo.DialInfo, error) {
	conf, err := getMongoConf(path)
	if err != nil {
		return nil, err
	}
	info := &mgo.DialInfo{
        Addrs:    []string{conf.Hostname},
        Timeout:  60 * time.Second,
        Database: conf.Database,
        Username: conf.Username,
        Password: conf.Password,
    }
	return info, nil
}