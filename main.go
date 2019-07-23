package main

import (
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

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/places", placesIndex).Methods("GET")
	// router.HandleFunc("/api/places", Store).Methods("POST")
	// router.HandleFunc("/api/places/{id}", Update).Methods("PATCH")
	// router.HandleFunc("/api/places/{id}", Delete).Methods("DELETE")

	log.Panic(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
