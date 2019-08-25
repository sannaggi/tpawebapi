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
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, Access-Control-Allow-Headers, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	fmt.Println(":" + port)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/places", getPlaces).Methods("GET")
	router.HandleFunc("/api/places/recommended", fetchRecommendedPlaces).Methods("GET")
	router.HandleFunc("/api/places/{id}", getPlace).Methods("GET")
	router.HandleFunc("/api/experiences", getExperiences).Methods("GET")
	router.HandleFunc("/api/experiences/recommended", fetchRecommendedExperiences).Methods("GET")
	router.HandleFunc("/api/experiences/{id}", getExperience).Methods("GET")
	router.HandleFunc("/api/search/{query}", getSearchResults).Methods("GET")
	router.HandleFunc("/api/experiences/search", searchExperienceByCategories).Methods("POST")
	router.HandleFunc("/api/experiences/s", fetchLimitedExperiences).Methods("POST")
	router.HandleFunc("/api/login/o", loginOauth2).Methods("POST")
	router.HandleFunc("/api/login/c", cookieLogin).Methods("POST")
	router.HandleFunc("/api/login/e", emailLogin).Methods("POST")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/users/", createNewUser).Methods("POST")
	router.HandleFunc("/api/users/check", checkNewUser).Methods("POST")
	router.HandleFunc("/api/wishlist/u/{id}", getUserWishlists).Methods("GET")
	router.HandleFunc("/api/wishlist/", addNewWishlist).Methods("POST")
	router.HandleFunc("/api/wishlist/d/{id}", removeFromWishlist).Methods("POST")
	router.HandleFunc("/api/wishlist/{id}", fetchWishlist).Methods("GET")
	router.HandleFunc("/api/wishlist/{id}", addToWishlist).Methods("POST")
	router.HandleFunc("/api/chat/u/{id}", getUserChat).Methods("GET")
	router.HandleFunc("/api/chat/{id}", changeChatStatus).Methods("POST")
	router.HandleFunc("/api/chat/{id}", getChat).Methods("GET")
	router.HandleFunc("/api/chat/m/{id}", addNewMessage).Methods("POST")
	router.HandleFunc("/api/chat", createNewChat).Methods("POST")

	log.Panic(http.ListenAndServe(":"+port, router))
}
