package collections

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wishlist struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Privacy     string             `json:"privacy"`
	UserID      string             `json:"userid" bson:"userid,omitempty"`
	Stays       []string           `json:"stays"`
	Experiences []string           `json:"experiences"`
}
