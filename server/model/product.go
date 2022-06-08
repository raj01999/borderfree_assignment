package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductField struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName   string             `json:"productname" bson:"productname"`
	ProductDetail string             `json:"productdetail" bson:"productdetail"`
	UserId        primitive.ObjectID `json:"userId" bson:"userId"`
}
