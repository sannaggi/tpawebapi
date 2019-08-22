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

	var chat c.Chat

	collection := client.Database("tpaweb").Collection("chat")

	id := params["id"]

	err := collection.FindOne(context.Background(), bson.M{"users": id}).Decode(&chat)
	CheckErr(err)

	json.NewEncoder(w).Encode(chat)
}
