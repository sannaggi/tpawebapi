package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Userid string             `json:"userid"`
	Name   string             `json:"name"`
	Date   string             `json:"date"`
	Total  float64            `json:"total"`
	Type   string             `json:"type"`
	Status string             `json:"status"`
}
