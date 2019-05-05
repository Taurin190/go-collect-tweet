package collect_tweet

type TwitterSecret struct {
	ConsumerKey string `JSON:"ConsumerKey"`
	ConsumerSecret string `JSON:"ConsumerSecret"`
	AccessToken string `JSON:"AccessToken"`
	AccessSecret string `JSON:"AccessSecret"`
}

type MongoConf struct {
	hostname string `JSON:"hostname"`
	database string `JSON:"database"`
	collection string `JSON:"collection"`
	username string `JSON:"username"`
	password string `JSON:"password"`
}

const TWEET_CONF string = "./conf/twitter.conf"
const MONGO_CONF string = "./conf/mongo.conf"
