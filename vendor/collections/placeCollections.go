package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Place struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name"`
	Price         int                `json:"price"`
	Rating        []ReviewDetail     `json:"rating"`
	Category      string             `json:"category"`
	HostID        string             `json:"hostid"`
	AverageRating float64            `json:"averagerating"`
	RatingCount   int                `json:"ratingcount"`
	Images        []string           `json:"images"`
	Guests        int                `json:"guests"`
	Bedrooms      int                `json:"bedrooms"`
	Beds          int                `json:"beds"`
	Baths         int                `json:"baths"`
}
