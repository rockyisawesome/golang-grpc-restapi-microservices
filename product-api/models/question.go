package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	type Question struct {
//		ID        primitive.ObjectID `bson:"_id"`
//		Title     string             `bson:"title" json:"title"`
//		Content   string             `bson:"email" json:"email"`
//		UserID    primitive.ObjectID `bson:"userId" json:"userId"`
//		CreatedAt time.Time          `bson:"created_at" json:"created_at"`
//	}
// type Question struct {
// 	ID        primitive.ObjectID `bson:"_id"`
// 	Title     string             `bson:"title"`
// 	Content   string             `bson:"content"`
// 	UserID    primitive.ObjectID `bson:"userId"`
// 	CreatedAt primitive.DateTime `bson:"created_at"`
// }

type Question struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	UserID    primitive.ObjectID `bson:"userId"`
	CreatedAt time.Time          `bson:"created_at"`
}
