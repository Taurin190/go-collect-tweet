package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"../../../go-collect-tweet"
	"gopkg.in/mgo.v2"
)

func main() {
	info := &mgo.DialInfo{
        Addrs:    []string{"localhost"},
        Timeout:  60 * time.Second,
        Database: "twitter",
        Username: "root",
        Password: "root",
    }

	tweets, err := collect_tweet.Exec()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	results, err := collect_tweet.InsertTweet(tweets, info)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	fmt.Println(results)
	os.Exit(1)
}