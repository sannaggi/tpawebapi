package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name"`
	Price         float64			 `json:"price"`
	Review        []ReviewDetail     `json:"rating"`
	Category      string             `json:"category"`
	HostID        string             `json:"hostid"`
	AverageRating float64            `json:"averagerating"`
	TotalRating   int                `json:"totalrating"`
	Location      string             `json:"location"`
	Duration      float64            `json:"duration"`
	Amenities     []Amenity          `json:"amenities"`
	HeaderImage   string             `json:"headerimage"`
	Story         []string           `json:"story"`
	Detail        string             `json:"detail"`
	Gallery       []string           `json:"gallery"`
	AboutHost     string             `json:"abouthost"`
	Requirement   []string           `json:"requirement"`
	Guests        int                `json:"guests"`
	Languages     []string           `json:"languages"`
}
