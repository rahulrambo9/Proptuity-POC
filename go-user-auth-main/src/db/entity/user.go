package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    int
	Name  string
	Email string
}

type ServiceKey struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	KeyType   string             `json:"key_type,omitempty"`
	KeyId     string             `json:"key_id,omitempty"`
	KeySecret string             `json:"key_secret,omitempty"`
	CreatedTS int64              `json:"createdTS,omitempty" bson:"createdTS,omitempty"`
}
