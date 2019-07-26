package main

import (
	c "collections"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func getExperiences(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")

	var experiences []c.Experience

	collection := client.Database("tpaweb").Collection("experience")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	CheckErr(err)
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var experience c.Experience
		cursor.Decode(&experience)
		experiences = append(experiences, experience)
	}
	json.NewEncoder(w).Encode(experiences)
}

func getExperience(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	var experience c.Experience

	collection := client.Database("tpaweb").Collection("experience")

	id, err := primitive.ObjectIDFromHex(params["id"])
	CheckErr(err)

	fmt.Println(id)

	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&experience)
	CheckErr(err)

	json.NewEncoder(w).Encode(experience)
}
