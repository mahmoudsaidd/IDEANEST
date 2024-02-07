package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo/pkg/database/mongodb"
)

func main() {
	// Connect to MongoDB
	client, err := mongodb.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	// Close the connection when the main function exits
	defer client.Disconnect(context.Background())

	// Your application logic goes here
}
