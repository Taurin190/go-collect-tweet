package service

import (
	"fmt"

	"../entity"
	"../outer"
	"../repository"
)

type TweetService interface {
	Exec() (string, error)
}

/*
GetTweetService : get tweet service instance of TweetService interface
*/
func GetTweetService(r repository.TweetRepository, o outer.TweetOuterAPI) TweetService {
	return &tweetService{
		Repository: r,
		OuterAPI:   o,
	}
}

type tweetService struct {
	Repository repository.TweetRepository
	OuterAPI   outer.TweetOuterAPI
	Tweets     []entity.Tweet
}

func (service *tweetService) Exec() (string, error) {
	executionResults := ""
	results, err := service.setTweetsFromOuterAPI()
	if err != nil {
		return "", err
	}
	executionResults = executionResults + results
	results, err = service.insertTweets()
	if err != nil {
		return "", err
	}
	executionResults = executionResults + results
	return executionResults, nil
}

func (service *tweetService) setTweetsFromOuterAPI() (string, error) {
	tweets, err := service.OuterAPI.GetTweets()
	if err != nil {
		return "", err
	}
	service.Tweets = tweets
	return "Success: Get Tweets from Twitter API\n", nil
}

func (service *tweetService) insertTweets() (string, error) {
	count, err := service.Repository.Insert(service.Tweets)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Success: Insert %d of Tweet\n", count), nil
}
