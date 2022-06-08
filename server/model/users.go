package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserField struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}
