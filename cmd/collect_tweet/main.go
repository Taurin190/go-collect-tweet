package main

import (
	"os"
	"fmt"
	"../../../go-collect-tweet"
)

func main() {
	tweet := collect_tweet.GetTweetData()
	fmt.Println(tweet)
	os.Exit(1)
}