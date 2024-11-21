package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func connectToMongo(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}
func insert(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a new document
	newUser := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
	}

	// Insert the document
	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}
	fmt.Printf("Inserted document with ID: %v\n", result.InsertedID)
}
func fetch(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("Failed to find documents: %v", err)
	}
	defer cursor.Close(ctx)

	fmt.Println("Fetched documents:")
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			log.Fatalf("Failed to decode document: %v", err)
		}
		fmt.Printf("Name: %s, Email: %s, Age: %d\n", user.Name, user.Email, user.Age)
	}
	if err := cursor.Err(); err != nil {
		log.Fatalf("Cursor error: %v", err)
	}
}

func fetchDataWithoutStructure(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch all documents without specifying a structure
	filter := bson.D{} // Empty filter fetches all documents
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("Failed to fetch documents: %v", err)
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document
	fmt.Println("Fetched documents:")
	for cursor.Next(ctx) {
		var rawDocument bson.M // Use bson.M to handle any structure
		if err := cursor.Decode(&rawDocument); err != nil {
			log.Fatalf("Failed to decode document: %v", err)
		}
		fmt.Printf("%+v\n", rawDocument) // Print the raw document as a map
	}

	if err := cursor.Err(); err != nil {
		log.Fatalf("Cursor error: %v", err)
	}
}
