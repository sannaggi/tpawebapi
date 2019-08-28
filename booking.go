package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"
)

func addNewBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	collection := client.Database("tpaweb").Collection("booking")

	var booking c.Booking

	json.NewDecoder(r.Body).Decode(&booking)

	_, err := collection.InsertOne(context.Background(), booking)
	CheckErr(err)
}
