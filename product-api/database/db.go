package database

import (
	"context"
	"product-api/models"
)

type Database interface {
	Connect(ctx context.Context)
	Disconnect(ctx context.Context)

	ListUsers(ctx context.Context) ([]*models.Users, error)
}
