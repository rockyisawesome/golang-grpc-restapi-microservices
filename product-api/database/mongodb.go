package database

import (
	"context"
	"fmt"
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
	// Ctx      *context.Context
}

// returning an instance og MongoDB
func NewMongoDB(cfg *configs.MongoDbConfig, lobbs *hclog.Logger) *MongoDB {
	return &MongoDB{
		Config: cfg,
		loggs:  lobbs,
		// Ctx:    ctx,
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

func (mango *MongoDB) GetAllUsers(ctx context.Context) ([]*models.Users, error) {
	// getting user collection
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

func (mango *MongoDB) GetAllQuestion(ctx context.Context) ([]*models.Question, error) {
	// getting user collection
	questionCollection := mango.Database.Collection("questions")

	cursor, err := questionCollection.Find(ctx, bson.M{})
	if err != nil {
		(*mango.loggs).Error("Not able to find any question or some error occured", err)
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

func (mango *MongoDB) GetAllReplies(ctx context.Context) ([]*models.Reply, error) {
	// getting user collection
	repliesCollection := mango.Database.Collection("replies")

	cursor, err := repliesCollection.Find(ctx, bson.M{})
	if err != nil {
		(*mango.loggs).Error("Not able to find any Replies or some error occured", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var replies []*models.Reply
	if err = cursor.All(ctx, &replies); err != nil {
		(*mango.loggs).Error("Cursor all not working", err)
		return nil, err
	}
	return replies, nil
}

func (mango *MongoDB) GetAllUserQuestions(ctx context.Context, objectId string) ([]*models.Question, error) {
	// getting user collection
	questionsCollection := mango.Database.Collection("questions")

	// getting mongo db object id
	obid, err := bson.ObjectIDFromHex(objectId)
	fmt.Println(obid)
	if err != nil {
		(*mango.loggs).Error("passed string is not able to be parsed to Object Id", err)
		return nil, err
	}

	cursor, err := questionsCollection.Find(ctx, bson.M{"userId": obid})
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

func (mango *MongoDB) GetAllQuestionReplies(ctx context.Context, objectId string) ([]*models.Reply, error) {
	// getting user collection
	repliesCollection := mango.Database.Collection("replies")

	// getting mongo db object id
	obid, err := bson.ObjectIDFromHex(objectId)
	fmt.Println(obid)
	if err != nil {
		(*mango.loggs).Error("passed string is not able to be parsed to Object Id", err)
		return nil, err
	}

	cursor, err := repliesCollection.Find(ctx, bson.M{"questionId": obid})
	if err != nil {
		(*mango.loggs).Error("Not able to find users", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var replies []*models.Reply
	if err = cursor.All(ctx, &replies); err != nil {
		(*mango.loggs).Error("Cursor all not working", err)
		return nil, err
	}
	return replies, nil
}

func (mango *MongoDB) GetUserProfileWithQuestion(ctx context.Context, userId string) (*models.UsersWithQuestion, error) {

	userCollection := mango.Database.Collection("users")

	// getting mongo db object id
	obid, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		(*mango.loggs).Error("passed string is not able to be parsed to Object Id", err)
		return nil, err
	}

	mongoSingleResult := userCollection.FindOne(ctx, bson.M{"_id": obid})
	// handling errors from the mongo single result
	if err := mongoSingleResult.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			(*mango.loggs).Error("No Document found with the given ObjectId", err)
		} else {
			(*mango.loggs).Error("Some other error", err)
		}
		return nil, err
	}

	var userModel *models.Users
	if err := mongoSingleResult.Decode(&userModel); err != nil {
		(*mango.loggs).Error("Error in decoding the single user result", err)
		return nil, err
	}

	// getting the questions of the users
	// can we not call the GetAllUserQuestions method
	questions, err := mango.GetAllUserQuestions(ctx, userId)
	if err != nil {
		(*mango.loggs).Error("No Questions Found for this user", err)
	}

	UsersWithQuestion := &models.UsersWithQuestion{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Role:      userModel.Role,
		Questions: questions,
	}

	return UsersWithQuestion, nil
}

func (mango *MongoDB) GetAllQuestionsAndReplies(ctx context.Context) ([]*models.QuestionWithReply, error) {

	allQuestions, err := mango.GetAllQuestion(ctx)
	if err != nil {
		(*mango.loggs).Error("Some error occured in fetching questions", err)
		return nil, err
	}

	var questionWithReply []*models.QuestionWithReply

	for _, question := range allQuestions {

		replies, err := mango.GetAllQuestionReplies(ctx, question.ID.Hex())
		var repliesWithSubReply []*models.ReplyWithSubReply
		if err != nil {
			(*mango.loggs).Error("Error in finding Question replies", err)
		} else {

			for _, repl := range replies {
				newReplyWithSubReply := &models.ReplyWithSubReply{
					ID:         repl.ID,
					Content:    repl.Content,
					QuestionId: repl.QuestionId,
					UserId:     repl.UserId,
				}
				subreply, err := mango.GetSubReplyFromReply(ctx, repl.ID.Hex())
				if err != nil {
					(*mango.loggs).Error("Error in finding Subreply replies", err)
				}

				newReplyWithSubReply.SubReplies = subreply
				repliesWithSubReply = append(repliesWithSubReply, newReplyWithSubReply)
			}
		}

		newQuesReply := &models.QuestionWithReply{
			ID:      question.ID,
			Title:   question.Title,
			Content: question.Content,
			UserID:  question.UserID,
			Replies: repliesWithSubReply,
		}
		questionWithReply = append(questionWithReply, newQuesReply)

	}

	return questionWithReply, nil

}

func (mango *MongoDB) GetSubReplyFromReply(ctx context.Context, objectId string) ([]*models.SubReply, error) {
	// getting user collection
	subReplyCollection := mango.Database.Collection("subreplies")

	// getting mongo db object id
	obid, err := bson.ObjectIDFromHex(objectId)
	fmt.Println(obid)
	if err != nil {
		(*mango.loggs).Error("passed string is not able to be parsed to Object Id", err)
		return nil, err
	}

	cursor, err := subReplyCollection.Find(ctx, bson.M{"replyId": obid})
	if err != nil {
		(*mango.loggs).Error("Not able to find users", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var subreplies []*models.SubReply
	if err = cursor.All(ctx, &subreplies); err != nil {
		(*mango.loggs).Error("Cursor all not working", err)
		return nil, err
	}
	return subreplies, nil
}
