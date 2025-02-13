package models

// type Users struct {
// 	ID        primitive.ObjectID `bson: "_id""`
// 	Name      string             `bson:"name"`
// 	Email     string             `bson:"email"`
// 	CreatedAt time.Time          `bson:"created_at"`
// }

// type Users struct {
// 	ID        primitive.ObjectID `bson:"id"`
// 	Name      string             `bson:"name"`
// 	Email     string             `bson:"email"`
// 	CreatedAt time.Time          `bson:"created_at"`
// }

type Users struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}
