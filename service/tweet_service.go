package service

type TweetService interface {
}

func GetTweetService() TweetService {
	return &tweetService{}
}

type tweetService struct {
}

func (service *tweetService) Insert() error {
	return nil
}
