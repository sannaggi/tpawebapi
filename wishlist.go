package main

import (
	c "collections"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type wishlistItem struct {
	ID      string `json:"id"`
	IsPlace bool   `json:"isplace"`
}

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

	wishlist.Stays = []string{}
	wishlist.Experiences = []string{}

	_, err := collection.InsertOne(context.Background(), wishlist)
	CheckErr(err)
}

func fetchWishlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	params := mux.Vars(r)

	collection := client.Database("tpaweb").Collection("wishlist")

	var wishlist c.Wishlist
	id, err := primitive.ObjectIDFromHex(params["id"])

	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&wishlist)
	CheckErr(err)

	json.NewEncoder(w).Encode(wishlist)
}

func addToWishlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	params := mux.Vars(r)
	wishlistID, err := primitive.ObjectIDFromHex(params["id"])

	var item wishlistItem
	json.NewDecoder(r.Body).Decode(&item)

	var field string
	if item.IsPlace == true {
		field = "stays"
	} else {
		field = "experiences"
	}

	collection := client.Database("tpaweb").Collection("wishlist")

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": wishlistID}, bson.M{"$push": bson.M{field: item.ID}})
	if err != nil {
		fmt.Println(err.Error())
	}
	CheckErr(err)
}
