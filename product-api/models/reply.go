package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Reply struct {
	ID         bson.ObjectID `bson:"_id" json:"_id"`
	Content    string        `bson:"content" json:"content"`
	QuestionId bson.ObjectID `bson:"questionId" json:"questionId"`
	UserId     bson.ObjectID `bson:"userId" json:"userId"`
	CreatedAt  time.Time     `bson:"created_at" json:"-"`
}

type ReplyWithSubReply struct {
	ID         bson.ObjectID `bson:"_id" json:"_id"`
	Content    string        `bson:"content" json:"content"`
	QuestionId bson.ObjectID `bson:"questionId" json:"questionId"`
	UserId     bson.ObjectID `bson:"userId" json:"userId"`
	SubReplies []*SubReply   `bson:"subreplies" json:"subreplies"`
}
