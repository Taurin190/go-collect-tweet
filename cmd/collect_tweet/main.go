package main

import (
	"os"
	"fmt"
	"log"
	"../../../go-collect-tweet"
)

func main() {
	tweets, err := collect_tweet.Exec()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	info, err := collect_tweet.GetMongoInfo()
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