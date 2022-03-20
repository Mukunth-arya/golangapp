package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName    string             `json:"Product,omitempty"`
	Servicecomment string             `json:"servicecomment,omitempty"`
	Qualiycomment  string             `json:"qualitycomment,omitempty"`
	Satisfied      bool               `json:"satisfied,omitempty"`
}
