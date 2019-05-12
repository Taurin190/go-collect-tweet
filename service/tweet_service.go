package service

import (
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
	Tweets     entity.Tweet
}

func (service *tweetService) Exec() (string, error) {
	executionResults := ""
	results, err := service.setTweetsFromOuterAPI()
	if err != nil {
		return executionResults, err
	}
	executionResults = executionResults + results
	results, err = service.setupRepository()
	if err != nil {
		return executionResults, err
	}
	executionResults = executionResults + results
	results, err = service.insertTweets()
	if err != nil {
		return executionResults, err
	}
	executionResults = executionResults + results
	return executionResults, nil
}

func (service *tweetService) setTweetsFromOuterAPI() (string, error) {
	return "", nil
}

func (service *tweetService) setupRepository() (string, error) {
	return "", nil
}

func (service *tweetService) insertTweets() (string, error) {
	return "", nil
}
