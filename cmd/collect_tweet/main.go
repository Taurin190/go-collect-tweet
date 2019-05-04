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
	fmt.Println(tweets)
	os.Exit(1)
}