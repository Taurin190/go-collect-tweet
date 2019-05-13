package entity

import (
	"github.com/ChimeraCoder/anaconda"
)

/*
Tweet structure to insert mongo db
*/
type Tweet struct {
	ID    int64
	IDStr string
	Text  string
	User
}

/*
User structure of twitter
*/
type User struct {
	ID                   int64
	IDStr                string
	Name                 string
	ScreenName           string
	Location             string
	URL                  string
	Description          string
	FollowersCount       int
	FriendsCount         int
	ProfileImageURLHTTPS string
}

/*
ConvertTweetsStruct : From anaconda.Tweet, convert tweet structure
*/
func ConvertTweetsStruct(anacondaTweets []anaconda.Tweet) []Tweet {
	var tweets []Tweet
	for _, anacondaTweet := range anacondaTweets {
		var tweet Tweet
		tweet.ID = anacondaTweet.Id
		tweet.IDStr = anacondaTweet.IdStr
		tweet.Text = anacondaTweet.Text

		var user User
		user.ID = anacondaTweet.User.Id
		user.IDStr = anacondaTweet.User.IdStr
		user.Name = anacondaTweet.User.Name
		user.ScreenName = anacondaTweet.User.ScreenName
		user.Location = anacondaTweet.User.Location
		user.URL = anacondaTweet.User.URL
		user.Description = anacondaTweet.User.Description
		user.FollowersCount = anacondaTweet.User.FollowersCount
		user.FriendsCount = anacondaTweet.User.FriendsCount
		user.ProfileImageURLHTTPS = anacondaTweet.User.ProfileBackgroundImageUrlHttps

		tweet.User = user

		tweets = append(tweets, tweet)
	}
	return tweets
}
