package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name"`
	Price         int                `json:"price"`
	Review        []ReviewDetail     `json:"rating"`
	Category      string             `json:"category"`
	HostID        string             `json:"hostid"`
	AverageRating float64            `json:"averagerating"`
	RatingCount   int                `json:"ratingcount"`
	Images        []string           `json:"images"`
}

type ReviewDetail struct {
	UserID string  `json:"userid"`
	Review string  `json:"review"`
	Rating float64 `json:"rating"`
}
