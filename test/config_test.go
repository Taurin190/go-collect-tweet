package test

import (
	"testing"

	"../config"
	"github.com/stretchr/testify/assert"
)

func Test_GetConfig_Normal(t *testing.T) {
	path := config.ConfigPath{
		MongoDBConfigPath: "../conf/mongo.conf",
		TwitterConfigPath: "../conf/twitter.conf.test",
	}
	config := config.GetConfig(path)
	assert.Equal(t, "AAAAAA", config.Twitter.ConsumerKey)
	assert.Equal(t, "BBBBBB", config.Twitter.ConsumerSecret)
	assert.Equal(t, "CCCCCC", config.Twitter.AccessToken)
	assert.Equal(t, "DDDDDD", config.Twitter.AccessSecret)
}
