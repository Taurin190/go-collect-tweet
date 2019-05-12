package main

import (
	"fmt"
	"log"
	"os"

	config "../../config"
	o "../../outer"
	r "../../repository"
	s "../../service"
)

func main() {
	path := config.ConfigPath{
		MongoDBConfigPath: "./conf/mongo.conf",
		TwitterConfigPath: "./conf/twitter.conf",
	}
	c := config.GetConfig(path)
	repository := r.GetTweetRepository(c)
	outerAPI := o.GetTweetOuterAPI(c)
	service := s.GetTweetService(repository, outerAPI)

	results, err := service.Exec()

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	fmt.Println(results)

	os.Exit(1)
}
