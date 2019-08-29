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

type bookingStatus struct {
	Status string `json:"status"`
}

type rating struct {
	Type   string  `json:"type"`
	Rating float64 `json:"rating"`
}

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

func getUserBookings(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var bookings []c.Booking

	collection := client.Database("tpaweb").Collection("booking")

	cursor, err := collection.Find(context.Background(), bson.M{"userid": id})
	CheckErr(err)

	for cursor.Next(context.TODO()) {
		var booking c.Booking
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	json.NewEncoder(w).Encode(bookings)
}

func changeBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	var status bookingStatus
	json.NewDecoder(r.Body).Decode(&status)

	collection := client.Database("tpaweb").Collection("booking")

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status.Status}})
	CheckErr(err)
}

func setRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	var rating rating
	json.NewDecoder(r.Body).Decode(&rating)

	collection := client.Database("tpaweb").Collection("booking")

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{rating.Type: rating.Rating}})
	CheckErr(err)
}
