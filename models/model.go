package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	CakeName    string             `json:"CakeName" `
	Cakeflavour string             `json:"cakeflavour" `
	TypeofCream string             `json:"typeofcream"`
	Toppings    string             `json:"toppings" `
	Shape       string             `json:"shapeofcake"`
	Satisfied   bool               `json:"satisfied"`
}
