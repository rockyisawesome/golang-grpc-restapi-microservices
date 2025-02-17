package database

import (
	"context"
	"product-api/models"
)

type Database interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	GetAllUsers(ctx context.Context) ([]*models.Users, error)
	GetAllQuestion(ctx context.Context) ([]*models.Question, error)
	GetUserProfileWithQuestion(ctx context.Context, userId string) (*models.UsersWithQuestion, error)
	GetAllUserQuestions(ctx context.Context, objectId string) ([]*models.Question, error)
	GetAllQuestionsAndReplies(ctx context.Context) ([]*models.QuestionWithReply, error)
	GetAllReplies(ctx context.Context) ([]*models.Reply, error)
	GetAllQuestionReplies(ctx context.Context, objectId string) ([]*models.Reply, error)
	GetSubReplyFromReply(ctx context.Context, objectId string) ([]*models.SubReply, error)
}
