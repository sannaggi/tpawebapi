package main

import (
	c "collections"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	j "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

type oauth struct {
	ID            string `json:"id"`
	Authenticator string `json:"authenticator"`
}

type regResponse struct {
	Result bool `json:"result"`
}

func generateToken(user c.User) string {
	token := j.New(j.SigningMethodHS256)
	claims := token.Claims.(j.MapClaims)
	claims["user"] = user

	tokenString, err := token.SignedString([]byte("nolep"))
	CheckErr(err)

	return tokenString
}

func loginOauth2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	var oauth oauth
	json.NewDecoder(r.Body).Decode(&oauth)

	var user c.User
	collection := client.Database("tpaweb").Collection("user")

	fmt.Println(oauth.ID)
	err := collection.FindOne(context.Background(), bson.M{oauth.Authenticator: oauth.ID}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(user)
	}

	tokenString := generateToken(user)

	json.NewEncoder(w).Encode(tokenString)
}

func checkNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	collection := client.Database("tpaweb").Collection("user")

	var user c.User
	var err error

	json.NewDecoder(r.Body).Decode(&user)
	if user.FacebookID != "" {
		err = collection.FindOne(context.Background(), bson.M{"facebookid": user.FacebookID}).Decode(&user)
	} else {
		err = collection.FindOne(context.Background(), bson.M{"googleid": user.GoogleID}).Decode(&user)
	}

	if err == nil {
		json.NewEncoder(w).Encode(false)
	} else if err.Error() == "mongo: no documents in result" {
		json.NewEncoder(w).Encode(true)
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	collection := client.Database("tpaweb").Collection("user")

	var user c.User
	var err error

	json.NewDecoder(r.Body).Decode(&user)

	user.Description = "Hi there! I'm using aivbnb"

	_, err = collection.InsertOne(context.Background(), user)
	CheckErr(err)
}
