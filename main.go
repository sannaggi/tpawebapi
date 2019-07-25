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
	// router.HandleFunc("/api/places", Store).Methods("POST")
	// router.HandleFunc("/api/places/{id}", Update).Methods("PATCH")
	// router.HandleFunc("/api/places/{id}", Delete).Methods("DELETE")

	log.Panic(http.ListenAndServe(":"+port, router))
}
