package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func getUserChat(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	var chats []c.Chat

	collection := client.Database("tpaweb").Collection("chat")

	id := params["id"]

	cursor, err := collection.Find(context.Background(), bson.M{"users": id})
	CheckErr(err)

	CheckErr(err)

	for cursor.Next(context.TODO()) {
		var chat c.Chat
		cursor.Decode(&chat)
		chats = append(chats, chat)
	}

	json.NewEncoder(w).Encode(chats)
}
