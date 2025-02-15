package models


import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)


// type Users struct {
// 	ID        primitive.ObjectID `bson: "_id""`
// 	Name      string             `bson:"name"`
// 	Email     string             `bson:"email"`
// 	CreatedAt time.Time          `bson:"created_at"`
// }


type Users struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	CreatedAt time.Time     `bson:"created_at" json:"-"`
}

type UserAPI struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
