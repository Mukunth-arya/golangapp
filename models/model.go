package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	ProductName    string             `json:"Product" `
	Servicecomment string             `json:"servicecomment" `
	Qualiycomment  string             `json:"qualitycomment"`
	Satisfied      bool               `json:"satisfied" `
}
