package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Actor struct {
	firstname string
	lastname  string
	awards    int16
}

func main() {
	//set client option
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(">> Connected to MongoDB!")

	collection := client.Database("dvdstore").Collection("actordetails")

	//setting up filter
	filter := bson.D{}
	//filter := bson.D{{"firstname", "James"}}

	var result Actor

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found a single actor: ", result)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(">> Disconnected from MongoDB!")
	}
}
