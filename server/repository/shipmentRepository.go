package repository

import (
	"context"
	"sync"

	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShipmentRepository struct {
	db *mongo.Collection
}

var onceSR sync.Once

func NewShipmentRepository(DB *Connection) *ShipmentRepository {
	var repository *ShipmentRepository
	onceSR.Do(func() {
		repository = new(ShipmentRepository)
		repository.db = DB.OpenCollection("logistics", "shipments")
	})
	return repository
}

func (store *ShipmentRepository) Insert(ctx context.Context, shipment *models.Shipment) (*mongo.InsertOneResult, error) {
	result, err := store.db.InsertOne(ctx, shipment)
	return result, err
}

func (store *ShipmentRepository) FindOne(ctx context.Context, id primitive.ObjectID) *mongo.SingleResult {
	result := store.db.FindOne(ctx, bson.M{"_id": id})
	return result
}

func (store *ShipmentRepository) FindAll(ctx context.Context) (*mongo.Cursor, error) {
	result, err := store.db.Find(ctx, bson.M{})
	return result, err
}

func (store *ShipmentRepository) UpdateStatus(ctx context.Context, shipmentId primitive.ObjectID, status string) (*mongo.UpdateResult, error) {
	result, err := store.db.UpdateByID(ctx, bson.M{"_id": shipmentId},
		bson.M{"$set": bson.M{"status": status}})
	return result, err
}
