package collect_tweet

import (
	"log"
	"github.com/ChimeraCoder/anaconda"
	"github.com/globalsign/mgo"
)

func UpsertTweet (tweets []anaconda.Tweet, mongo_host string) (string, error){
	session, err := mgo.Dial("mongodb://localhost/twitter")
	defer session.Close()
	db := session.DB("test")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	for _, tweet := range tweets {
        db.C("tweet").Insert(tweet)
    }
	return "", err
}