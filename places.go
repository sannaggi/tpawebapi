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

func getPlaces(w http.ResponseWriter, r *http.Request) {
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")

	var places []c.Place

	collection := client.Database("tpaweb").Collection("places")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	CheckErr(err)
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var place c.Place
		cursor.Decode(&place)
		places = append(places, place)
	}
	json.NewEncoder(w).Encode(places)
}

func getPlace(w http.ResponseWriter, r *http.Request) {
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	var place c.Place

	collection := client.Database("tpaweb").Collection("places")

	id, err := primitive.ObjectIDFromHex(params["id"])
	CheckErr(err)

	fmt.Println(id)

	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&place)
	CheckErr(err)

	json.NewEncoder(w).Encode(place)
}
