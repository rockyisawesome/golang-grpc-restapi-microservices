package models


import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Users struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Role      string        `bson:"role" json:"role"`
	CreatedAt time.Time     `bson:"created_at" json:"-"`
}

type UsersWithQuestion struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Role      string        `bson:"role" json:"role"`
	Questions []*Question   `bson:"questions" json:"questions"`
}
