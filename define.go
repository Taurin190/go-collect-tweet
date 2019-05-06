package collect_tweet

type TwitterSecret struct {
	ConsumerKey string `JSON:"ConsumerKey"`
	ConsumerSecret string `JSON:"ConsumerSecret"`
	AccessToken string `JSON:"AccessToken"`
	AccessSecret string `JSON:"AccessSecret"`
}

type MongoConf struct {
	Hostname string `JSON:"hostname"`
	Database string `JSON:"database"`
	Collection string `JSON:"collection"`
	Username string `JSON:"username"`
	Password string `JSON:"password"`
}

const TWEET_CONF string = "./conf/twitter.conf"
const MONGO_CONF string = "./conf/mongo.conf"
