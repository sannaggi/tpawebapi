package main

import (
	"context"
	"fmt"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbHandler struct {
	ConnectionString string
}

func (db *dbHandler) connect() *mongo.Client {
	db.ConnectionString = "mongodb+srv://sannaggi:dealbion11@tpaweb-jhprm.mongodb.net/test?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(db.ConnectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	fmt.Printf("%T", client)

	CheckErr(err)

	err = client.Ping(context.TODO(), nil)
	CheckErr(err)

	fmt.Println("Connected to MongoDB!")

	return client
}

// func (db *dbHandler) Query(command string) {
// 	// conn := db.connect()
// }
