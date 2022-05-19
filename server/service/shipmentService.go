package service

import (
	"context"
	"net/http"
	"sync"
	"time"

	"server/models"
	"server/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShipmentService struct {
	// inventoryRepository *repository.InventoryRepository
	shipmentRepository  *repository.ShipmentRepository
	inventoryRepository *repository.InventoryRepository
	LocationService     *LocationService
	validator           *validator.Validate
}

var onceShipment sync.Once

// Create an instance of the Shipment service
func NewShipmentService(shipRepository *repository.ShipmentRepository, invRepository *repository.InventoryRepository, locationService *LocationService) *ShipmentService {
	var instance *ShipmentService
	onceShipment.Do(func() {
		instance = new(ShipmentService)
		instance.shipmentRepository = shipRepository
		instance.inventoryRepository = invRepository
		instance.LocationService = locationService
		instance.validator = validator.New()
	})

	return instance
}

// Create new Shipment item
func (service *ShipmentService) CreateShipment(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var shipment models.Shipment
	if err := c.BindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"create Shipment error": err.Error()})
		return
	}

	validationErr := service.validator.Struct(shipment)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	shipment.Id = primitive.NewObjectID()

	result, err := service.shipmentRepository.Insert(ctx, &shipment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"create Shipment error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// List all Shipment items
func (service *ShipmentService) GetShipments(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var shipment []bson.M

	cursor, err := service.shipmentRepository.FindAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"get Shipment error": err.Error()})
		return
	}

	if err = cursor.All(ctx, &shipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"get Shipment error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipment)
}

// Get Shipment item by id
func (service *ShipmentService) GetShipmentById(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	shipmentID, _ := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	var shipment bson.M

	if err := service.shipmentRepository.FindOne(ctx, shipmentID).Decode(&shipment); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"get item Shipment error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipment)
}

// Deliver Shipment and update inventory
func (service *ShipmentService) DeliverShipment(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	ShipmentId, _ := primitive.ObjectIDFromHex(c.Params.ByName("id"))

	result, err := service.shipmentRepository.UpdateStatus(ctx, ShipmentId, "shipped")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"deliver shipment error": err.Error()})
		return
	}

	shipment, err := service.getShipment(ShipmentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"deliver shipment error": err.Error()})
		return
	}

	for _, item := range shipment.Items {
		inventory, err := service.GetInventory(item.InventoryId)
		if err == nil {
			service.inventoryRepository.DecreaseStock(ctx, inventory.Id, item.Stock)
		}
	}

	c.JSON(http.StatusOK, result.ModifiedCount)
}

// Receive Shipment and update inventory
func (service *ShipmentService) ReceiveShipment(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	ShipmentId, _ := primitive.ObjectIDFromHex(c.Params.ByName("id"))

	result, err := service.shipmentRepository.UpdateStatus(ctx, ShipmentId, "delivered")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"recieve shipment error": err.Error()})
		return
	}

	shipment, err := service.getShipment(ShipmentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"recieve shipment error": err.Error()})
		return
	}

	for _, item := range shipment.Items {
		inventory, err := service.GetInventoryAtLocation(item.Name, shipment.City, shipment.Warehouse)
		if err == nil {
			service.inventoryRepository.IncreaseStock(ctx, inventory.Id, item.Stock)
		} else {
			inventory := models.Inventory{
				Id:          primitive.NewObjectID(),
				Name:        item.Name,
				City:        shipment.City,
				Warehouse:   shipment.Warehouse,
				CostPerUnit: item.CostPerUnit,
				Stock:       item.Stock,
			}
			service.inventoryRepository.Insert(ctx, &inventory)
		}
	}

	c.JSON(http.StatusOK, result.ModifiedCount)
}

func (service *ShipmentService) getShipment(id primitive.ObjectID) (models.Shipment, error) {
	var shipment models.Shipment
	err := service.shipmentRepository.FindOne(context.TODO(), id).Decode(&shipment)
	return shipment, err
}

func (service *ShipmentService) GetInventory(id primitive.ObjectID) (models.Inventory, error) {
	var inventory models.Inventory
	err := service.inventoryRepository.FindOne(context.TODO(), id).Decode(&inventory)
	return inventory, err
}

func (service *ShipmentService) GetInventoryAtLocation(name string, city string, warehouse string) (models.Inventory, error) {
	var inventory models.Inventory
	err := service.inventoryRepository.FindAtLocation(context.TODO(), name, city, warehouse).Decode(&inventory)
	return inventory, err
}
