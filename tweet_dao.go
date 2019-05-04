package collect_tweet

import (
	"log"
	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/mgo.v2"
)

func InsertTweet (tweets []anaconda.Tweet, info *mgo.DialInfo) (string, error){
	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	db := session.DB(info.Database)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	for _, tweet := range tweets {
        db.C("tweet").Insert(tweet)
    }
	return "", err
}