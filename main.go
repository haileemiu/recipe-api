package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recipe struct {
	Name string
}

func main() {
	fmt.Println("main")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("recipe-api").Collection("recipes")

	a := Recipe{"Spagetti"}
	b := Recipe{"Tacos"}
	c := Recipe{"Chicken Fingers"}

	insertResult, err := collection.InsertOne(context.TODO(), a)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	recipes := []interface{}{b, c}

	insertManyResult, err := collection.InsertMany(context.TODO(), recipes)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// err = client.Disconnect(context.TODO())

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connection to MongoDB closed.")
}
