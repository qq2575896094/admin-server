package utils

import (
	"context"
	"fmt"
	"github.com/qq2575896094/admin-server/conf"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var RedisClient *redis.Client // RedisClient Redis缓存客户端
var MongoClient *mongo.Client // MongoClient MongoDB缓存客户端

var RedisContext = context.Background()

// InitRedisClient 初始化Redis
func initRedisClient() {
	fmt.Println("Connecting to Redis...")
	redisConfig := conf.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Username: redisConfig.Username,
		Password: redisConfig.Password,
		DB:       redisConfig.DbName,
	})

	_, err := client.Ping(RedisContext).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Redis!")
	RedisClient = client
}

func getMongoUri() string {
	mongoConf := conf.Config.Mongodb
	uri := "mongodb://" + mongoConf.Username + ":" + mongoConf.Password + "@" + mongoConf.Host + ":" + mongoConf.Port
	return uri
}

// InitMongoClient 初始化mongodb
func initMongoClient() {
	fmt.Println("Connecting to MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoUri()))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	MongoClient = client
}

func InitServer() {
	initMongoClient()
	initRedisClient()
}

func CloseServer(ctx context.Context) {
	_ = MongoClient.Disconnect(ctx)
	_ = RedisClient.Close()
}
