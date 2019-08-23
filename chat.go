package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type status struct {
	Status string `json:"status"`
	Value  bool   `json:"value"`
}

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

	for cursor.Next(context.TODO()) {
		var chat c.Chat
		cursor.Decode(&chat)
		chats = append(chats, chat)
	}

	json.NewEncoder(w).Encode(chats)
}

func changeChatStatus(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")

	var s status

	collection := client.Database("tpaweb").Collection("chat")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	json.NewDecoder(r.Body).Decode(&s)
	CheckErr(err)

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{s.Status: s.Value}})
	CheckErr(err)
}
