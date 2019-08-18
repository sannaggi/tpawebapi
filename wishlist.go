package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func getUserWishlists(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("content-type", "application/json")
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	params := mux.Vars(r)
	id := params["id"]

	collection := client.Database("tpaweb").Collection("wishlist")

	var wishlists []c.Wishlist

	cursor, err := collection.Find(context.TODO(), bson.M{"userid": id})
	CheckErr(err)
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var wishlist c.Wishlist
		cursor.Decode(&wishlist)
		wishlists = append(wishlists, wishlist)
	}

	json.NewEncoder(w).Encode(wishlists)
}

func addNewWishlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	collection := client.Database("tpaweb").Collection("wishlist")

	var wishlist c.Wishlist

	json.NewDecoder(r.Body).Decode(&wishlist)

	_, err := collection.InsertOne(context.Background(), wishlist)
	CheckErr(err)
}
