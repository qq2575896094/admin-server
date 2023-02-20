package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id       primitive.ObjectID `bson:"_id"`
	Title    string
	Type     string
	Tag      string
	FavCount int
	Author   string
}
