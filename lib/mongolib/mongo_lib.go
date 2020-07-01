package mongolib

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	client *mongo.Client
)

func NewClient() {
	addr := viper.GetString("mongo.addr")
	password := viper.GetString("mongo.password")
	db := viper.GetString("mongo.db")
	username := viper.GetString("mongo.username")
	url := "mongodb://" + username + ":" + password + "@" + addr + "/" + db

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println("初始化mongoDb", url)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	client = mongoClient
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(fmt.Errorf("Fatal error mongolib connect: %s \n", err))
	}
}

func GetConn(coll string) *mongo.Collection {
	collection := client.Database("prim").Collection(coll)
	return collection
}

func InsertOne(coll *mongo.Collection, m interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := coll.InsertOne(ctx, m)
	if err != nil {
		fmt.Println("mongodb 添加数据异常", err)
	}
	id := res.InsertedID
	return id, err
}

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}
func FindOne(coll *mongo.Collection, filter interface{}, info interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := coll.FindOne(ctx, filter).Decode(info)
	return err
}
