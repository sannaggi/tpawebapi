package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	fmt.Println(":" + port)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/places", getPlaces).Methods("GET")
	router.HandleFunc("/api/places/{id}", getPlace).Methods("GET")
	router.HandleFunc("/api/experiences", getExperiences).Methods("GET")
	router.HandleFunc("/api/experiences/{id}", getExperience).Methods("GET")
	router.HandleFunc("/api/search/{query}", getSearchResults).Methods("GET")

	log.Panic(http.ListenAndServe(":"+port, router))

	// var query string
	// fmt.Scanln(&query)

	// var arr []kv

	// getPlacesFromDb()
	// getExperiencesFromDb()
	// fmt.Println("aa")

	// for _, place := range getPlacesFromDb() {
	// 	value := strings.Index(place.Name, query)
	// 	if value == -1 {
	// 		continue
	// 	}
	// 	arr = append(arr, kv{
	// 		data{
	// 			place.ID,
	// 			"place",
	// 			place.Name,
	// 			place.Price,
	// 			place.Category,
	// 			place.AverageRating,
	// 			place.RatingCount,
	// 			place.Images[0],
	// 		},
	// 		value,
	// 	})
	// }
	// for _, experience := range getExperiencesFromDb() {
	// 	value := strings.Index(experience.Name, query)
	// 	if value == -1 {
	// 		continue
	// 	}
	// 	arr = append(arr, kv{
	// 		data{
	// 			experience.ID,
	// 			"experience",
	// 			experience.Name,
	// 			experience.Price,
	// 			experience.Category,
	// 			experience.AverageRating,
	// 			experience.TotalRating,
	// 			experience.HeaderImage,
	// 		},
	// 		value,
	// 	})
	// }

	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i].Value < arr[j].Value
	// })

	// for _, info := range arr {
	// 	fmt.Printf("%s, %d\n", info.Object.Name, info.Value)
	// }
}
