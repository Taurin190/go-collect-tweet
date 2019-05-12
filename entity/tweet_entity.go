package entity

/*
Tweet structure to insert mongo db
*/
type Tweet struct {
	ID    int64
	IDStr string
	Text  string
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
