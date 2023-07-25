package dao

import (
	"context"
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserCollection() {
	userCollection = utils.MongoClient.Database("appdb").Collection("users")
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
	return userCollection.FindOne(ctx, bson.D{{"_id", id}}).Decode(userInfo)
}
