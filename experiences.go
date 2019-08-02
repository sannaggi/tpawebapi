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

type category struct {
	Guests     int      `json:"guests"`
	Lowerprice int      `json:"lowerprice"`
	Upperprice int      `json:"upperprice"`
	Languages  []string `json:"languages"`
}

func searchExperienceByCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	var cat category
	json.NewDecoder(r.Body).Decode(&cat)

	var experiences []c.Experience
	var languages []string
	for _, language := range cat.Languages {
		languages = append(languages, language)
	}
	collection := client.Database("tpaweb").Collection("experience")
	cursor, err := collection.Find(context.Background(), bson.M{"guests": bson.M{"$gte": cat.Guests}, "price": bson.M{"$gte": cat.Lowerprice, "$lte": cat.Upperprice}, "languages": bson.M{"$in": languages}})
	CheckErr(err)

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
