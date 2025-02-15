package database

import (
	"context"
	configs "product-api/configurations"
	"product-api/models"

	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Config   *configs.MongoDbConfig
	loggs    *hclog.Logger
}

// returning an instance og MongoDB
func NewMongoDB(cfg *configs.MongoDbConfig, lobbs *hclog.Logger) *MongoDB {
	return &MongoDB{
		Config: cfg,
		loggs:  lobbs,
	}
}

func (mango *MongoDB) Connect(ctx context.Context) error {

	client, err := mongo.Connect(options.Client().ApplyURI(mango.Config.MongoURI))
	if err != nil {
		(*mango.loggs).Error("Error connecting to Mongo DB", "Error", err)
		return err
	}

	// ping the database to verify connections
	err = client.Ping(ctx, nil)
	if err != nil {
		(*mango.loggs).Error("Pinging Database but no response", "Error", err)
		return err
	}
	(*mango.loggs).Info("Database is up and active", "Error", err)

	mango.Client = client
	mango.Database = client.Database(mango.Config.DBName)
	(*mango.loggs).Info("Connected to Database", "DB", mango.Config.DBName)

	return nil
}

// Disconnect implements the Database interface
func (mango *MongoDB) Disconnect(ctx context.Context) error {
	return mango.Client.Disconnect(ctx)
}

// get all the user
func (mango *MongoDB) ListUsers(ctx context.Context) ([]*models.Users, error) {
	userCollection := mango.Database.Collection("users")
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		(*mango.loggs).Error("Not able to find users", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.Users
	if err = cursor.All(ctx, &users); err != nil {
		(*mango.loggs).Error("Cursor all not working", err)
		return nil, err
	}
	return users, nil
}

func (mango *MongoDB) ListQuestions(ctx context.Context) ([]*models.Question, error) {
	questionCollection := mango.Database.Collection("questions")
	// userCollection := mango.database.Collection("users")
	// // find one user
	// cu, err := userCollection.Find(ctx, bson.M{})
	// if err != nil {
	// 	(*mango.loggs).Error("Not able to find anything in users collection")
	// }
	// defer cu.Close(ctx)
	// var users []*models.Users
	// if err = cu.All(ctx, &users); err != nil {
	// 	(*mango.loggs).Info("Error in User Collections", err)
	// }

	// if len(users) > 0 {
	// 	(*mango.loggs).Info("There are items in the list", users[0].ID.String(), users[0].ID.Hex())
	// } else {
	// 	(*mango.loggs).Info("No Item in User Collection", err)
	// }

	// uid, err := primitive.ObjectIDFromHex("67ad798199d610f8824d7943")
	// (*mango.loggs).Info("_id", uid.String())
	// if err != nil {
	// 	(*mango.loggs).Error("Error occured here")
	// }
	// cursor, err := questionCollection.Find(ctx, bson.M{"_id": users[0].ID})
	cursor, err := questionCollection.Find(ctx, bson.M{})
	if err != nil {
		(*mango.loggs).Error("Not able to find users", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var questions []*models.Question
	if err = cursor.All(ctx, &questions); err != nil {
		(*mango.loggs).Error("Cursor all not working", err)
		return nil, err
	}
	return questions, nil
}
