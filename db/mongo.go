package db

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

type Book struct {
	Id       primitive.ObjectID `bson:"_id"`
	Title    string
	Type     string
	Tag      string
	FavCount int
	Author   string
}

var Client *mongo.Client

func InitServer() {
	dsn := "mongodb://" + viper.GetString("mongodb.username") + ":" + viper.GetString("mongodb.password") + "@" + viper.GetString("mongodb.host") + ":" + viper.GetString("mongodb.port") + "/?authMechanism=SCRAM-SHA-1"
	var err error

	Client, err = mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		log.Panic("mongodb error: ", err.Error())
	}

	if err := Client.Connect(getContext()); err != nil {
		log.Panic("mongodb connect error: ", err.Error())
	}

	if err := Client.Ping(getContext(), nil); err != nil {
		log.Panic("mongodb ping error: ", err)
	} else {
		log.Println("mongodb server is successful~~")
	}
}
