package repository

import (
	"context"
	"shopify-challenge/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type inventoryRepository struct {
	db *mongo.Collection
}

func NewInventoryRepository(DB *mongo.Database) models.InventoryRepository {
	return &inventoryRepository{
		db: DB.Collection("inventory"),
	}
}

func (store *inventoryRepository) Insert(inventory *models.Inventory) (*models.Inventory, error) {
	result, err := store.db.InsertOne(context.TODO(), inventory)
	return result, err
}

func (store *inventoryRepository) FindOne(id primitive.ObjectID) (*models.Inventory, error) {
	var result models.Inventory
	err := store.db.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&result)
	return result, err
}

func (store *inventoryRepository) FindAll() ([]models.Inventory, error) {
	var result []models.Inventory
	err := store.db.Find(context.TODO(), bson.D{}).Decode(&result)
	return result, err
}

func (store *inventoryRepository) Delete(inventoryID primitive.ObjectID) error {
	result, err := store.db.DeleteOne(context.TODO(), bson.D{{"_id", inventoryID}})
	return err
}

func (store *inventoryRepository) Update(inventory *models.Inventory) error {
	result, err := store.db.ReplaceOne(context.TODO(), bson.D{{"_id", inventory.id}}, inventory)
	return err
}
