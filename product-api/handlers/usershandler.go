package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"product-api/database"
	"product-api/models"

	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserHandler struct {
	loggs  *hclog.Logger
	userdb *database.UsersDB
	ctx    *context.Context
}

type UserWithQuestion struct {
}

func NewUserHandler(l *hclog.Logger, userdb *database.UsersDB, contuxt *context.Context) *UserHandler {
	return &UserHandler{
		loggs:  l,
		userdb: userdb,
		ctx:    contuxt,
	}
}

func (userHandle *UserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	(*userHandle.loggs).Info("A call made to Hello GetMethod")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		(*userHandle.loggs).Error("Some error occured", err)
		http.Error(rw, "some error", http.StatusBadGateway)
		return
	}

	(*userHandle.loggs).Info("request body", "body", string(d))
	data, err := (*userHandle).getAllUsers()
	if err != nil {
		(*userHandle.loggs).Error("Some issue occured in data fetching")
		http.Error(rw, "Dekho error aa gayi bhai", http.StatusBadRequest)
		return
	}

	(*userHandle.loggs).Info("Data fetched properly")

	json.NewEncoder(rw).Encode(data)
	return
}

func (userHandle *UserHandler) getAllUsers() ([]*models.Users, error) {
	// userCollection := mango.database.Collection("users")
	(*userHandle.loggs).Info("I am in getAllUsers")

	userCollection := userHandle.userdb.MongodbClient.Database.Collection("users")

	(*userHandle.loggs).Info("Got the userCollection from mongodb client")

	cursor, err := userCollection.Find(*userHandle.ctx, bson.M{"name": "Alice Johnson"})

	if err != nil {
		(*userHandle.loggs).Error("Not able to find users", err)
		return nil, err
	}
	(*userHandle.loggs).Info("userCollection Find is working")
	defer cursor.Close(*userHandle.ctx)

	var users []*models.Users
	if err = cursor.All(*userHandle.ctx, &users); err != nil {
		(*userHandle.loggs).Error("Cursor all not working", err)
		return nil, err
	}

	(*userHandle.loggs).Info("lets see i reach here or not")
	return users, nil

}

// func (userHandle *UserHandler) getAllUsersWithQuestions() error {

// }
