package constants

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// 系统用户默认常量
var (
	UserAvatarDefault = "https://avatars.githubusercontent.com/u/32?v=4" // 用户头像
	UserGenderDefault = "男"                                              // 用户性别
	UserTypeDefault   = "user"                                           // 用户类型
)

// db常量
var (
	RedisClient  *redis.Client // RedisClient Redis缓存客户端
	MongoClient  *mongo.Client // MongoClient MongoDB缓存客户端
	RedisContext = context.Background()
)
