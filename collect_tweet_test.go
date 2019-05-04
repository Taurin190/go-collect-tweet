package collect_tweet

import (
	"testing"
	"fmt"
)

func Test_LoadConfig(t *testing.T) {
	s, err := LoadConfig("./conf/twitter.conf.test")
	if err != nil {
		t.Errorf("LoadConfig($directory_path) has some error to load")
	}
	var in, expected string
	in = "./conf/twitter.conf.test"
	expected = "ConsumerKey"
	fmt.Println(s.ConsumerKey)
	if s.ConsumerKey != expected {
		t.Errorf("LoadConfig(%s) = TwitterSecret.%s, want %s", in, s, expected)
	}
	expected = "ConsumerSecret"
	if s.ConsumerSecret != expected {
		t.Errorf("LoadConfig(%s) = TwitterSecret.%s, want %s", in, s, expected)
	}
	expected = "AccessToken"
	if s.AccessToken != expected {
		t.Errorf("LoadConfig(%s) = TwitterSecret.%s, want %s", in, s, expected)
	}
	expected = "AccessToken"
	if s.AccessToken != expected {
		t.Errorf("LoadConfig(%s) = TwitterSecret.%s, want %s", in, s, expected)
	}
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