package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	Id          primitive.ObjectID `json:"id" bson:"id"`
	Name        string             `json:"name"`
	Stock       int                `json:"stock"`
	CostPerUnit float32            `json:"costPerUnit"`
	City        string             `json:"city"`
	Warehouse   string             `json:"warehouse"`
}
