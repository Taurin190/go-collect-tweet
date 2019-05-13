package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"../entity"
	o "../outer"
	r "../repository"
	s "../service"
)

func Test_Exec_Normal(t *testing.T) {
	repository := getDummyTweetRepository()
	outerAPI := getDummyTweetOuterAPI()
	service := s.GetTweetService(repository, outerAPI)

	results, err := service.Exec()
	assert.Equal(t, "Success: Get Tweets from Twitter API\nSuccess: Setup Repository\nSuccess: Insert 0 of Tweet\n", results)
	assert.Nil(t, err)
}

func getDummyTweetRepository() r.TweetRepository {
	return &dummyTweetRepository{}
}

func getDummyTweetOuterAPI() o.TweetOuterAPI {
	return &dummyTweetOuterAPI{}
}

type dummyTweetRepository struct {
}

func (repository *dummyTweetRepository) SetupRepository() error {
	return nil
}

func (repository *dummyTweetRepository) Insert(tweets []entity.Tweet) (int, error) {
	return 0, nil
}

type dummyTweetOuterAPI struct {
}

func (outerAPI *dummyTweetOuterAPI) GetTweets() ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	return tweets, nil
}
