package main

import (
	"context"
	"fmt"
	"log"

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
	james := Actor{"James", "Roger", 9}
	mark := Actor{"Mark", "Brown", 0}
	mili := Actor{"Mili", "Ford", 11}

	actors := []interface{}{mark, mili}

	insertResult, err := collection.InsertOne(context.TODO(), james)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a new actor: ", insertResult)

	insertManyResult, err := collection.InsertMany(context.TODO(), actors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted new actors: ", insertManyResult.InsertedIDs)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(">> Disconnected from MongoDB!")
	}
}
