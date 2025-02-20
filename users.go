package main

import (
	c "collections"
	"context"
	"encoding/json"
	"net/http"

	j "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type oauth struct {
	ID            string `json:"id"`
	Authenticator string `json:"authenticator"`
}

type regResponse struct {
	Result bool `json:"result"`
}

type loginRequest struct {
	ID string `json:"id"`
}

type emailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func generateToken(user c.User) string {
	token := j.New(j.SigningMethodHS256)
	claims := token.Claims.(j.MapClaims)
	claims["user"] = user

	tokenString, err := token.SignedString([]byte("nolep"))
	CheckErr(err)

	return tokenString
}

func emailLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	var req emailLoginRequest
	json.NewDecoder(r.Body).Decode(&req)
	var user c.User
	collection := client.Database("tpaweb").Collection("user")

	err := collection.FindOne(context.Background(), bson.M{"email": req.Email, "password": req.Password}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(nil)
		return
	}

	tokenString := generateToken(user)

	json.NewEncoder(w).Encode(tokenString)
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

	err := collection.FindOne(context.Background(), bson.M{oauth.Authenticator: oauth.ID}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(nil)
		return
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
	} else if user.GoogleID != "" {
		err = collection.FindOne(context.Background(), bson.M{"googleid": user.GoogleID}).Decode(&user)
	} else {
		err = collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&user)
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

func cookieLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

	var reqID loginRequest
	json.NewDecoder(r.Body).Decode(&reqID)

	var user c.User
	collection := client.Database("tpaweb").Collection("user")

	id, err := primitive.ObjectIDFromHex(reqID.ID)
	CheckErr(err)

	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(nil)
		return
	}

	tokenString := generateToken(user)

	json.NewEncoder(w).Encode(tokenString)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	var user c.User

	collection := client.Database("tpaweb").Collection("user")

	id, err := primitive.ObjectIDFromHex(params["id"])
	CheckErr(err)

	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	CheckErr(err)
	json.NewEncoder(w).Encode(user)
}
