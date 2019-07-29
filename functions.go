package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type data struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type          string             `json:"type"`
	Name          string             `json:"name"`
	Price         int                `json:"price"`
	Category      string             `json:"category"`
	AverageRating float64            `json:"averagerating"`
	RatingCount   int                `json:"ratingcount"`
	Images        string             `json:"image"`
}

type kv struct {
	Object data
	Value  int
}

func getSearchResults(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	params := mux.Vars(r)
	query := params["query"]

	var results []data
	results = append(results, searchPlaceByName(query)...)
	results = append(results, searchExperienceByName(query)...)

	sort.Slice(results, func(i, j int) bool {
		return strings.Index(strings.ToLower(results[i].Name), strings.ToLower(query)) < strings.Index(strings.ToLower(results[j].Name), strings.ToLower(query))
	})

	if len(results) > 5 {
		results = results[0:5]
	}

	json.NewEncoder(w).Encode(results)
}
