package main

import (
	"os"
	"fmt"
	"../../../go-collect-tweet"
)

func main() {
	tweet := collect_tweet.Exec()
	fmt.Println(tweet)
	os.Exit(1)
}