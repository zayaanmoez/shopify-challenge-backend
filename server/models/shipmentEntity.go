package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShipmentItem struct {
	InventoryId primitive.ObjectID `json:"inventoryId" bson:"inventoryId"`
	Name        string             `json:"name"`
	CostPerUnit float32            `json:"costPerUnit"`
	Stock       int                `json:"stock"`
}

type Shipment struct {
	Id        primitive.ObjectID `json:"id" bson:"id"`
	Label     string             `json:"label"`
	Items     []ShipmentItem     `json:"items"`
	City      string             `json:"city"`
	Warehouse string             `json:"warehouse"`
	Status    string             `json:"status"`
}
