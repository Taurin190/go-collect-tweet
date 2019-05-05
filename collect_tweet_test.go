package collect_tweet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_GetTwitterSecret_Normal(t *testing.T) {
	s, err := GetTwitterSecret("./conf/twitter.conf.test")
	if err != nil {
		t.Errorf("LoadConfig($directory_path) has some error to load")
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


func testLoadConfig(t *testing.T, in string, expected TwitterSecret) {
	// s := LoadConfig("./secret/twitter.conf.test")
	// if s != expected {
	// 	t.Errorf("LoadConfig(%s) = %s, want %s", in, s, expected)
	// }
}

func TestGetTweetData(t *testing.T) {

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