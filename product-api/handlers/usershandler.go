package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"product-api/database"
	"github.com/hashicorp/go-hclog"

)

type UserHandler struct {
	loggs  *hclog.Logger
	userdb *database.UsersDB
	ctx    *context.Context
}


func NewUserHandler(l *hclog.Logger, userdb *database.UsersDB, contuxt *context.Context) *UserHandler {

	return &UserHandler{
		loggs:  l,
		userdb: userdb,
		ctx:    contuxt,
	}
}


func (userHandle *UserHandler) GetAllUsers(rw http.ResponseWriter, r *http.Request) {

	(*userHandle.loggs).Info("A call made to Hello GetMethod")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		(*userHandle.loggs).Error("Some error occured", err)
		http.Error(rw, "some error", http.StatusBadGateway)
		return
	}

	(*userHandle.loggs).Info("request body", "body", string(d))

	data, err := userHandle.userdb.DbClient.GetAllUsers(*userHandle.ctx)

	if err != nil {
		(*userHandle.loggs).Error("Some issue occured in data fetching")
		http.Error(rw, "Dekho error aa gayi bhai", http.StatusBadRequest)
		return
	}

	(*userHandle.loggs).Info("Data fetched properly")

	json.NewEncoder(rw).Encode(data)
}

func (userHandle *UserHandler) GetAllUserQuestions(rw http.ResponseWriter, r *http.Request) {
	(*userHandle.loggs).Info("A call made to Hello GetMethod")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		(*userHandle.loggs).Error("Some error occured", err)
		http.Error(rw, "some error", http.StatusBadGateway)
		return
	}

	(*userHandle.loggs).Info("request body", "body", string(d))
	data, err := userHandle.userdb.DbClient.GetAllUserQuestions(*userHandle.ctx, "67aeda2f99d610f8824d794e")
	if err != nil {
		(*userHandle.loggs).Error("Some issue occured in data fetching")
		http.Error(rw, "Dekho error aa gayi bhai", http.StatusBadRequest)
		return
	}

	(*userHandle.loggs).Info("Data fetched properly")

	json.NewEncoder(rw).Encode(data)
}

func (userHandle *UserHandler) GetUserProfileWithQuestion(rw http.ResponseWriter, r *http.Request) {
	(*userHandle.loggs).Info("A call made to Hello GetMethod")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		(*userHandle.loggs).Error("Some error occured", err)
		http.Error(rw, "some error", http.StatusBadGateway)
		return
	}

	(*userHandle.loggs).Info("request body", "body", string(d))
	data, err := userHandle.userdb.DbClient.GetUserProfileWithQuestion(*userHandle.ctx, "67aeda2f99d610f8824d794e")
	if err != nil {
		(*userHandle.loggs).Error("Some issue occured in data fetching")
		http.Error(rw, "Dekho error aa gayi bhai", http.StatusBadRequest)
		return
	}

	(*userHandle.loggs).Info("Data fetched properly")

	json.NewEncoder(rw).Encode(data)
}

func (userHandle *UserHandler) GetAllQuestionsAndReplies(rw http.ResponseWriter, r *http.Request) {
	(*userHandle.loggs).Info("A call made to Hello GetMethod")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		(*userHandle.loggs).Error("Some error occured", err)
		http.Error(rw, "some error", http.StatusBadGateway)
		return
	}

	(*userHandle.loggs).Info("request body", "body", string(d))
	data, err := userHandle.userdb.DbClient.GetAllQuestionsAndReplies(*userHandle.ctx)
	if err != nil {
		(*userHandle.loggs).Error("Some issue occured in data fetching")
		http.Error(rw, "Dekho error aa gayi bhai", http.StatusBadRequest)
		return
	}

	(*userHandle.loggs).Info("Data fetched properly")

	json.NewEncoder(rw).Encode(data)
}

