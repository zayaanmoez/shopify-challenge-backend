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

type InventoryService struct {
	inventoryRepository *repository.InventoryRepository
	LocationService     *LocationService
	validator           *validator.Validate
}

var onceInventory sync.Once

// Create an instance of the inventory service
func NewInventoryService(repository *repository.InventoryRepository, locationService *LocationService) *InventoryService {
	var instance *InventoryService
	onceInventory.Do(func() {
		instance = new(InventoryService)
		instance.inventoryRepository = repository
		instance.LocationService = locationService
		instance.validator = validator.New()
	})

	return instance
}

// Create new inventory item
func (service *InventoryService) AddInventory(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var inventory models.Inventory

	if err := c.BindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"create inventory error": err.Error()})
		return
	}

	validationErr := service.validator.Struct(inventory)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	inventory.Id = primitive.NewObjectID()

	result, err := service.inventoryRepository.Insert(ctx, &inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"create inventory error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// List all inventory items
func (service *InventoryService) GetInventory(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var inventory []bson.M

	cursor, err := service.inventoryRepository.FindAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"get inventory error": err.Error()})
		return
	}

	if err = cursor.All(ctx, &inventory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"get inventory error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// Get inventory item by id
func (service *InventoryService) GetInventoryById(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	inventoryID, _ := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	var inventory bson.M

	if err := service.inventoryRepository.FindOne(ctx, inventoryID).Decode(&inventory); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"get item inventory error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// Delete inventory item by id
func (service *InventoryService) DeleteInventory(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	inventoryID, _ := primitive.ObjectIDFromHex(c.Params.ByName("id"))

	result, err := service.inventoryRepository.Delete(ctx, inventoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"delete error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Update inventory item by id
func (service *InventoryService) UpdateInventory(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	inventoryId, _ := primitive.ObjectIDFromHex(c.Params.ByName("id"))

	var inventory models.Inventory
	if err := c.BindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"update inventory error": err.Error()})
		return
	}

	validationErr := service.validator.Struct(inventory)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"update inventory error": validationErr.Error()})
		return
	}

	result, err := service.inventoryRepository.Update(ctx, inventoryId, &inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"update inventory error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result.ModifiedCount)
}
