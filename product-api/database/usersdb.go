package database

import "github.com/hashicorp/go-hclog"

// get the MongoDB
type UsersDB struct {
	loggs    *hclog.Logger
	DbClient Database
}

func NewUsersDB(l *hclog.Logger, dbclient Database) *UsersDB {
	return &UsersDB{
		loggs:    l,
		DbClient: dbclient,
	}
}
