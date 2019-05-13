package test

import (
	"fmt"
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

func Test_Exec_ErrorSetupRepository(t *testing.T) {
	repository := getDummyTweetRepositoryErrorSetup()
	outerAPI := getDummyTweetOuterAPI()
	service := s.GetTweetService(repository, outerAPI)

	results, err := service.Exec()
	assert.Equal(t, "", results)
	assert.Equal(t, "Error: Setup error", err.Error())
}

func Test_Exec_ErrorInsert(t *testing.T) {
	repository := getDummyTweetRepositoryInsert()
	outerAPI := getDummyTweetOuterAPI()
	service := s.GetTweetService(repository, outerAPI)

	results, err := service.Exec()
	assert.Equal(t, "", results)
	assert.Equal(t, "Error: Insert error", err.Error())
}

func Test_Exec_ErrorGetTweets(t *testing.T) {
	repository := getDummyTweetRepository()
	outerAPI := getDummyTweetOuterAPIError()
	service := s.GetTweetService(repository, outerAPI)

	results, err := service.Exec()
	assert.Equal(t, "", results)
	assert.Equal(t, "Error: GetTweets error", err.Error())
}

func getDummyTweetRepository() r.TweetRepository {
	return &dummyTweetRepository{}
}

type dummyTweetRepository struct {
}

func (repository *dummyTweetRepository) SetupRepository() error {
	return nil
}

func (repository *dummyTweetRepository) Insert(tweets []entity.Tweet) (int, error) {
	return 0, nil
}

func getDummyTweetOuterAPI() o.TweetOuterAPI {
	return &dummyTweetOuterAPI{}
}

type dummyTweetOuterAPI struct {
}

func (outerAPI *dummyTweetOuterAPI) GetTweets() ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	return tweets, nil
}

func getDummyTweetRepositoryErrorSetup() r.TweetRepository {
	return &dummyTweetRepositoryErrorSetup{}
}

type dummyTweetRepositoryErrorSetup struct {
}

func (repository *dummyTweetRepositoryErrorSetup) SetupRepository() error {
	return fmt.Errorf("Error: %s", "Setup error")
}

func (repository *dummyTweetRepositoryErrorSetup) Insert(tweets []entity.Tweet) (int, error) {
	return 0, nil
}

func getDummyTweetRepositoryInsert() r.TweetRepository {
	return &dummyTweetRepositoryErrorInsert{}
}

type dummyTweetRepositoryErrorInsert struct {
}

func (repository *dummyTweetRepositoryErrorInsert) SetupRepository() error {
	return nil
}

func (repository *dummyTweetRepositoryErrorInsert) Insert(tweets []entity.Tweet) (int, error) {
	return 0, fmt.Errorf("Error: %s", "Insert error")
}

func getDummyTweetOuterAPIError() o.TweetOuterAPI {
	return &dummyTweetOuterAPIError{}
}

type dummyTweetOuterAPIError struct {
}

func (outerAPI *dummyTweetOuterAPIError) GetTweets() ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	return tweets, fmt.Errorf("Error: %s", "GetTweets error")
}
