package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo/options"

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

func getChat(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	var chat c.Chat

	collection := client.Database("tpaweb").Collection("chat")

	id, err := primitive.ObjectIDFromHex(params["id"])
	CheckErr(err)

	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&chat)
	CheckErr(err)

	json.NewEncoder(w).Encode(chat)
}

func addNewMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	CheckErr(err)

	collection := client.Database("tpaweb").Collection("chat")

	var message c.Message

	json.NewDecoder(r.Body).Decode(&message)

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$push": bson.M{"messages": message}})
	CheckErr(err)
}

func getSpecificChat(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")

	var cu []string
	var chat c.Chat

	json.NewDecoder(r.Body).Decode(&cu)

	collection := client.Database("tpaweb").Collection("chat")

	err := collection.FindOne(context.Background(), bson.M{"users": bson.M{"$all": cu}}).Decode(&chat)
	CheckErr(err)

	json.NewEncoder(w).Encode(chat)
}

func createNewChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	collection := client.Database("tpaweb").Collection("chat")

	var chat c.Chat
	json.NewDecoder(r.Body).Decode(&chat)
	chat.Messages = []c.Message{}

	_, err := collection.UpdateOne(context.Background(), bson.M{"users": bson.M{"$all": []bson.M{bson.M{"$elemMatch": bson.M{"$eq": chat.Users[0]}}, bson.M{"$elemMatch": bson.M{"$eq": chat.Users[1]}}}}}, bson.M{"$set": chat}, options.Update().SetUpsert(true))
	CheckErr(err)
}
