package main

import (
	"context"
	"fmt"
)

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func getDb() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:password@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	client, err := getDb()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(client.Database("db1").Collection("user1").InsertOne(context.TODO(), User{
	//	Id:   2,
	//	Name: "name3",
	//}))
	fmt.Println(client.Database("db1").Collection("fw").BulkWrite(context.TODO(), []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(),
	}))
}

func saveOne(ctx context.Context, data interface{}) error {
	return nil
}

func saveMany(ctx context.Context, data interface{}) error {
	return nil
}
