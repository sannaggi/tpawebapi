package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func placesIndex(w http.ResponseWriter, r *http.Request) {
	client := new(dbHandler).connect()
	// defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")

	var places []c.Place

	collection := client.Database("tpaweb").Collection("places")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	CheckErr(err)
	// defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var place c.Place
		cursor.Decode(&place)
		places = append(places, place)
	}
	// fmt.Printf("%+v", places)
	json.NewEncoder(w).Encode(places)

}
