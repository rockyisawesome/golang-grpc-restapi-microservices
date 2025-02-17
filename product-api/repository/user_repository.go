package repository

// import (
// 	"product-api/database"
// 	"product-api/models"

// 	"github.com/hashicorp/go-hclog"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// type UserRepository struct {
// 	Userdb *database.UsersDB
// 	loggs  *hclog.Logger
// }

// func NewUserRepository(userdb *database.UsersDB, lobbs *hclog.Logger) *UserRepository {
// 	return &UserRepository{
// 		loggs:  lobbs,
// 		Userdb: userdb,
// 	}
// }

// func (ur *UserRepository) GetAllUsers() ([]*models.Users, error) {
// 	// getting user collection
// 	ctx := *&ur.Userdb.DbClient
// 	userCollection := ur.Userdb.MongodbClient.Database.Collection("users")

// 	cursor, err := userCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		(*ur.loggs).Error("Not able to find users", err)
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var users []*models.Users
// 	if err = cursor.All(ctx, &users); err != nil {
// 		(*ur.loggs).Error("Cursor all not working", err)
// 		return nil, err
// 	}
// 	return users, nil
// }
