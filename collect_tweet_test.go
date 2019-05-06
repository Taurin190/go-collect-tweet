package collect_tweet

import (
	"fmt"
	"time"
	"testing"
	"gopkg.in/mgo.v2"
	"github.com/stretchr/testify/assert"
)

func Test_GetTwitterSecret_Normal(t *testing.T) {
	s, err := GetTwitterSecret("./conf/twitter.conf.test")
	if err != nil {
		t.Errorf("GetTwitterSecret($path) has some error to load")
	}
	assert.Equal(t, "AAAAAA", s.ConsumerKey)
	assert.Equal(t, "BBBBBB", s.ConsumerSecret)
	assert.Equal(t, "CCCCCC", s.AccessToken)
	assert.Equal(t, "DDDDDD", s.AccessSecret)
}

func Test_GetTwitterSecret_Abnormal_InvalidFileName(t *testing.T) {
	var expectedTwitterSecret TwitterSecret
	s, err := GetTwitterSecret("./conf/twitter.conf.test.notfound")
	assert.Equal(t, expectedTwitterSecret, s)
	assert.NotNil(t, err)
	assert.Equal(t, "open ./conf/twitter.conf.test.notfound: no such file or directory", err.Error())
}

func Test_GetTwitterSecret_Abnormal_InvalidFileFormat(t *testing.T) {
	var expectedTwitterSecret TwitterSecret
	s, err := GetTwitterSecret("./conf/twitter.conf.test.invalid")
	assert.Equal(t, expectedTwitterSecret, s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid character 'a' looking for beginning of value", err.Error())
}

func Test_GetMongoInfo_Normal(t *testing.T) {
	s, err := GetMongoInfo("./conf/mongo.conf.test")
	if err != nil {
		t.Errorf("GetMongoInfo($path) has some error to load")
	}
	fmt.Println(s)
	assert.Equal(t, []string{"localhost"}, s.Addrs)
	assert.Equal(t, time.Duration(60000000000), s.Timeout)
	assert.Equal(t, "twitter", s.Database)
	assert.Equal(t, "root", s.Username)
	assert.Equal(t, "root", s.Password)
}

func Test_GetMongoInfo_Abnormal_InvalidFileName(t *testing.T) {
	var expectedMongoInfo *mgo.DialInfo
	s, err := GetMongoInfo("./conf/mongo.conf.test.notfound")
	assert.Equal(t, expectedMongoInfo, s)
	assert.NotNil(t, err)
	assert.Equal(t, "open ./conf/mongo.conf.test.notfound: no such file or directory", err.Error())
}

func Test_GetMongoInfo_Abnormal_InvalidFileFormat(t *testing.T) {
	var expectedMongoInfo *mgo.DialInfo
	s, err := GetMongoInfo("./conf/mongo.conf.test.invalid")
	assert.Equal(t, expectedMongoInfo, s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid character 'a' looking for beginning of value", err.Error())
}

func testTweetData(t *testing.T, in, expected string) {
	// s := TweetData("aa")
	// // if err != nil {
	// // 	log.Fatalf("err: %s", err)
	// // }
	// if s != expected {
	// 	t.Errorf("LoadConfig(%s) = %s, want %s", in, s, expected)
	// }
}