package database

import "github.com/hashicorp/go-hclog"

// get the MongoDB
type UsersDB struct {
	loggs         *hclog.Logger
	MongodbClient *MongoDB
}

func NewUsersDB(l *hclog.Logger, mongodbClient *MongoDB) *UsersDB {
	return &UsersDB{
		loggs:         l,
		MongodbClient: mongodbClient,
	}
}
