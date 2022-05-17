package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Variant struct {
	label string
	stock int
}

type Inventory struct {
	id primitive.ObjectID
	name string
	quantity int
	value float32
	costPerUnit float32
	warehouse string
	variants []Variant
}

type InventoryService interface {
	CreateInventory(inventory *Inventory) (*Inventory, error)
	ListAllInventory() ([]Inventory, error)
	DeleteInventory(inventoryID primitive.ObjectID) error
	UpdateInventory(inventory *Inventory) error
}

type InventoryRepository interface {
	Insert(inventory *Inventory) (*Inventory, error)
	FindOne(id primitive.ObjectID) (*Inventory, error)
	FindAll() ([]Inventory, error)
	Delete(inventoryID primitive.ObjectID) error
	Update(inventory *Inventory) error
}