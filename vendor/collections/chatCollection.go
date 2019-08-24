package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Users    []string           `json:"users"`
	Host     string             `json:"host"`
	Starred  bool               `json:"starred"`
	Archived bool               `json:"archived"`
	Unread   bool               `json:"unread"`
	Status   string             `json:"status"`
	Price    float64            `json:"price"`
	Messages []Message          `json:"messages"`
}

type Message struct {
	SenderID string `json:"senderid"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Time     string `json:"time"`
}
