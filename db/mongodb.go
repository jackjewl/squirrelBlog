package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodbClient *mongo.Client

//初始化mongodb连接
func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return
	}

	mongodbClient = client

}

//关闭mongodb连接
func CloseMongoDB(conn *mongo.Client) {
	if err := conn.Disconnect(ctx); err != nil {
		panic(err)
	}
}
