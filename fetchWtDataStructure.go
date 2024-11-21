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

func main() {
	// Replace with your MongoDB URI
	uri := "mongodb://localhost:27017"
	client := connectToMongo(uri)

	// Access the database and collection
	database := client.Database("testdb")         // Replace "testdb" with your database name
	collection := database.Collection("users")   // Replace "users" with your collection name

	// Fetch and print data without knowing the structure
	fetchDataWithoutStructure(collection)

	// Disconnect from MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
	fmt.Println("Disconnected from MongoDB")
}
