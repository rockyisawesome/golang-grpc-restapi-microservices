package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type SubReply struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Content   string        `bson:"content" json:"content"`
	ReplyId   bson.ObjectID `bson:"replyId" json:"replyId"`
	CreatedAt time.Time     `bson:"created_at" json:"-"`
}
