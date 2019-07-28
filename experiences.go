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

func getExperiencesFromDb() []c.Experience {
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

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
	return experiences
}

func getExperiences(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(getExperiencesFromDb())
}

func searchExperienceByName(query string) []data {
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	collection := client.Database("tpaweb").Collection("experience")
	cursor, err := collection.Find(context.Background(), bson.M{"name": bson.M{"$regex": query, "$options": "i"}})
	CheckErr(err)
	defer cursor.Close(context.TODO())

	var datas []data

	for cursor.Next(context.TODO()) {
		var experience c.Experience
		cursor.Decode(&experience)
		datas = append(datas, data{
			experience.ID,
			"experience",
			experience.Name,
			experience.Price,
			experience.Category,
			experience.AverageRating,
			experience.TotalRating,
			experience.HeaderImage,
		})
	}

	return datas
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
