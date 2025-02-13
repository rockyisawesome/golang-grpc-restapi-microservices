package repository

import (
	"github.com/hashicorp/go-hclog"
	mongo "go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	loggs      hclog.Logger
}

func NewUserRepository(db *mongo.Database, lobbs hclog.Logger) *UserRepository {
	return &UserRepository{
		loggs: lobbs,
		collection: db.Collection("users"),
	}
}
