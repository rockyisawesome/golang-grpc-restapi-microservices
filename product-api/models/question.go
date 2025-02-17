package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Question struct {
	ID        bson.ObjectID `bson:"_id"`
	Title     string        `bson:"title" json:"title"`
	Content   string        `bson:"content" json:"content"`
	UserID    bson.ObjectID `bson:"userId" json:"userId"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}

type QuestionWithReply struct {
	ID      bson.ObjectID        `bson:"_id"`
	Title   string               `bson:"title" json:"title"`
	Content string               `bson:"content" json:"content"`
	UserID  bson.ObjectID        `bson:"userId" json:"userId"`
	Replies []*ReplyWithSubReply `bson:"replies" json:"replies"`
}
