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
	Price         float64            `json:"price"`
	Category      string             `json:"category"`
	AverageRating float64            `json:"averagerating"`
	RatingCount   int                `json:"ratingcount"`
	Images        string             `json:"image"`
	HostID        string             `json:"hostid"`
}

type kv struct {
	Object data
	Value  int
}

func getSearchResults(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	params := mux.Vars(r)
	query := params["query"]

	// val, err := redisClient.Get("src-" + query).Result()
	// if err == nil {
	// 	var redisData data
	// 	fmt.Println([]byte(val))
	// 	err = json.Unmarshal([]byte(val), &redisData)
	// 	// CheckErr(err)
	// 	// fmt.Println(val)
	// 	// fmt.Println()
	// 	json.NewEncoder(w).Encode(redisData)
	// 	return
	// }

	var results []data
	results = append(results, searchPlaceByName(query)...)
	results = append(results, searchExperienceByName(query)...)

	sort.Slice(results, func(i, j int) bool {
		return strings.Index(strings.ToLower(results[i].Name), strings.ToLower(query)) < strings.Index(strings.ToLower(results[j].Name), strings.ToLower(query))
	})

	if len(results) > 5 {
		results = results[0:5]
	}

	// fmt.Println(results)
	// fmt.Println()
	// out, err := json.Marshal(results)
	// CheckErr(err)
	// fmt.Println(out)
	// fmt.Println()
	// fmt.Println(string(out))
	// fmt.Println()

	// err = redisClient.Set("src-"+query, out, time.Minute*30).Err()
	// CheckErr(err)

	// fmt.Println(results)

	json.NewEncoder(w).Encode(results)
}
