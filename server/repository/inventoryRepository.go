package repository

import (
	"context"
	"sync"

	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryRepository struct {
	db *mongo.Collection
}

var onceIR sync.Once

func NewInventoryRepository(DB *Connection) *InventoryRepository {
	var repository *InventoryRepository
	onceIR.Do(func() {
		repository = new(InventoryRepository)
		repository.db = DB.OpenCollection("logistics", "inventory")
	})
	return repository
}

func (store *InventoryRepository) Insert(ctx context.Context, inventory *models.Inventory) (*mongo.InsertOneResult, error) {
	result, err := store.db.InsertOne(ctx, inventory)
	return result, err
}

func (store InventoryRepository) FindOne(ctx context.Context, id primitive.ObjectID) *mongo.SingleResult {
	result := store.db.FindOne(ctx, bson.M{"_id": id})
	return result
}

func (store *InventoryRepository) FindAtLocation(ctx context.Context, name string, city string, warehouse string) *mongo.SingleResult {
	result := store.db.FindOne(ctx, bson.M{"name": name, "city": city, "warehouse": warehouse})
	return result
}

func (store *InventoryRepository) FindAll(ctx context.Context) (*mongo.Cursor, error) {
	result, err := store.db.Find(ctx, bson.M{})
	return result, err
}

func (store *InventoryRepository) Delete(ctx context.Context, inventoryID primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := store.db.DeleteOne(ctx, bson.M{"_id": inventoryID})
	return result, err
}

func (store *InventoryRepository) Update(ctx context.Context, inventoryId primitive.ObjectID, inventory *models.Inventory) (*mongo.UpdateResult, error) {
	result, err := store.db.UpdateOne(ctx, bson.M{"_id": inventoryId},
		bson.M{
			"name":        inventory.Name,
			"stock":       inventory.Stock,
			"costPerUnit": inventory.CostPerUnit,
			"city":        inventory.City,
			"warehouse":   inventory.Warehouse,
		})
	return result, err
}

func (store *InventoryRepository) IncreaseStock(ctx context.Context, inventoryId primitive.ObjectID, stock int) (*mongo.UpdateResult, error) {
	result, err := store.db.UpdateOne(ctx, bson.M{"_id": inventoryId},
		bson.M{"$inc": bson.M{
			"stock": stock,
		}})
	return result, err
}

func (store *InventoryRepository) DecreaseStock(ctx context.Context, inventoryId primitive.ObjectID, stock int) (*mongo.UpdateResult, error) {
	result, err := store.db.UpdateOne(ctx, bson.M{"_id": inventoryId},
		bson.M{"$dec": bson.M{
			"stock": stock,
		}})
	return result, err
}
