package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	ctx := context.TODO()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("testdb").Collection("users")
	_, err = collection.InsertOne(ctx, User{Name: "Gaurav", Age: 40})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Document inserted.")

	// Find
	var result User
	err = collection.FindOne(ctx, bson.M{"name": "Gaurav"}).Decode(&result)
	if err == nil {
		fmt.Println("Found:", result.Name)
	}
}
