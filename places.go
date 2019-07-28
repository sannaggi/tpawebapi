package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func getPlacesFromDb() []c.Place {
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

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

	return places
}

func getPlaces(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(getPlacesFromDb())
}

func searchPlaceByName(query string) []data {
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	collection := client.Database("tpaweb").Collection("places")
	cursor, err := collection.Find(context.Background(), bson.M{"name": bson.M{"$regex": query, "$options": "i"}})
	CheckErr(err)
	defer cursor.Close(context.TODO())

	var datas []data

	for cursor.Next(context.TODO()) {
		var place c.Place
		cursor.Decode(&place)
		datas = append(datas, data{
			place.ID,
			"place",
			place.Name,
			place.Price,
			place.Category,
			place.AverageRating,
			place.RatingCount,
			place.Images[0],
		})
	}

	return datas
}

func getPlace(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	var place c.Place

	collection := client.Database("tpaweb").Collection("places")

	id, err := primitive.ObjectIDFromHex(params["id"])
	CheckErr(err)

	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&place)
	CheckErr(err)

	json.NewEncoder(w).Encode(place)
}
