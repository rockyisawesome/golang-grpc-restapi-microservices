package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type FollowRelation struct {
	ID          bson.ObjectID `bson:"_id"`
	FollowerID  bson.ObjectID `bson:"follower_id"`
	FollowingID bson.ObjectID `bson:"following_id"`
	CreatedAt   time.Time     `bson:"created_at"`
}
