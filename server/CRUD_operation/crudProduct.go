package crud

import (
	"context"
	"server/connection"
	"server/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertProduct(product model.ProductField) (*mongo.InsertOneResult, error) {
	inserted, err := connection.Product.InsertOne(context.Background(), product)
	return inserted, err
}

func FindProduct(userId primitive.ObjectID) (*mongo.Cursor, error) {
	data, err := connection.Product.Find(context.Background(), bson.M{"userId": userId})
	return data, err
}

func FindOneProduct(id primitive.ObjectID) *mongo.SingleResult {
	data := connection.Product.FindOne(context.Background(), bson.M{"_id": id})
	return data
}

func DeleteProduct(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	deleted, err := connection.Product.DeleteOne(context.Background(), bson.M{"_id": id})
	return deleted, err
}

func UpdateProduct(id primitive.ObjectID, data model.ProductField) (*mongo.UpdateResult, error) {
	updated, err := connection.Product.UpdateOne(context.Background(), bson.M{"_id": id}, data)

	return updated, err
}
