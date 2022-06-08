package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductField struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName   string             `json:"productname,omitempty" bson:"productname,omitempty"`
	ProductDetail string             `json:"productdetail,omitempty" bson:"productdetail,omitempty"`
	UserId        primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
}
