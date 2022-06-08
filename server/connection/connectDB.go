package connection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:admin@cluster0.y05mh.mongodb.net/borderfree?retryWrites=true&w=majority"
const dbName = "borderfree"

var User *mongo.Collection

var Product *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)

	User = db.Collection("user")
	Product = db.Collection("product")

}
