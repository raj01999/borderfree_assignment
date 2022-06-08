package crud

import (
	"context"
	"server/connection"
	"server/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(user model.UserField) (*mongo.InsertOneResult, error) {
	inserted, err := connection.User.InsertOne(context.Background(), user)
	return inserted, err
}

func FindUser(email string) *mongo.SingleResult {
	data := connection.User.FindOne(context.Background(), bson.M{"username": email})
	return data
}
