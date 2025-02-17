package repository

import (
	"product-api/database"

	"github.com/hashicorp/go-hclog"
)

type QuestionRepository struct {
	Userdb *database.UsersDB
	loggs  *hclog.Logger
}

func NewUserRepository(userdb *database.UsersDB, lobbs *hclog.Logger) *UserRepository {
	return &UserRepository{
		loggs:  lobbs,
		Userdb: userdb,
	}
}
