package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type token struct {
	ID            string `json:"id"`
	Authenticator string `json:"authenticator"`
}

func loginOauth2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	var token token
	json.NewDecoder(r.Body).Decode(&token)

	var user c.User
	collection := client.Database("tpaweb").Collection("user")

	err := collection.FindOne(context.Background(), bson.M{token.Authenticator: token.ID}).Decode(&user)
	CheckErr(err)

	json.NewEncoder(w).Encode(user)
}
