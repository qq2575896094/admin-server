package dao

import (
	"context"
	"errors"
	"github.com/qq2575896094/admin-server/constants"
	"github.com/qq2575896094/admin-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserCollection() {
	userCollection = constants.MongoClient.Database("appdb").Collection("users")
}

//func SetUsernameUniqueIndexes() {
//
//}

// AddUser 添加用户
func AddUser(ctx context.Context, user *models.UserRegisterParams) (*mongo.InsertOneResult, error) {
	return userCollection.InsertOne(ctx, user)
}

// GetUserById 通过id获取用户信息
func GetUserById(ctx context.Context, id any, userInfo *models.UserInfo) error {
	switch id.(type) {
	case string:
		objectID, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return err
		}
		return userCollection.FindOne(ctx, bson.D{{"_id", objectID}}).Decode(userInfo)
	case primitive.ObjectID:
		return userCollection.FindOne(ctx, bson.D{{"_id", id}}).Decode(userInfo)
	default:
		return errors.New("非法数据: userId类型错误")
	}
}

// GetUserByName 通过用户名获取用户信息
func GetUserByName(ctx context.Context, username string, userInfo *models.UserInfo) error {
	return userCollection.FindOne(ctx, bson.D{{"username", username}}).Decode(userInfo)
}
