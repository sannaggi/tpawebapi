package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	GoogleID          string             `json:"googleid"`
	FacebookID        string             `json:"facebookid"`
	Email             string             `json:"email"`
	FirstName         string             `json:"firstname"`
	LastName          string             `json:"lastname"`
	ProfileImage      string             `json:"profileimage"`
	Gender            string             `json:"gender"`
	PhoneNumber       string             `json:"phonenumber"`
	PreferredLanguage string             `json:"preferredlanguage"`
	PreferredCurrency string             `json:"preferredcurrency"`
	Description       string             `json:"description"`
	SpokenLanguage    []string           `json:"spokenlanguage"`
	Review            []BasicReview      `json:"review"`
	ResponseRate      float64            `json:"responserate"`
	ResponseTime      int32              `json:"responsetime"`
	// UserDetail        UserDetail         `json:"userdetail"`
}

type UserDetail struct {
	GoogleID          string        `json:"googleid"`
	FacebookID        string        `json:"facebookid"`
	Email             string        `json:"email"`
	FirstName         string        `json:"firstname"`
	LastName          string        `json:"lastname"`
	ProfileImage      string        `json:"profileimage"`
	Gender            string        `json:"gender"`
	PhoneNumber       string        `json:"phonenumber"`
	PreferredLanguage string        `json:"preferredlanguage"`
	PreferredCurrency string        `json:"preferredcurrency"`
	Description       string        `json:"description"`
	SpokenLanguage    []string      `json:"spokenlanguage"`
	Review            []BasicReview `json:"review"`
	ResponseRate      float64       `json:"responserate"`
	ResponseTime      int32         `json:"responsetime"`
}
