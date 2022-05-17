package service

import (
	"shopify-challenge/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

var once sync.Once

type inventoryService struct {
	inventoryRepository models.InventoryRepository
}

var instance *inventoryService

func NewInventoryService(repository models.InventoryRepository) models.InventoryService {
	once.Do(func() {
		instance = &inventoryService{
			inventoryRepository: repository,
		}
	})

	return instance
}

func (service *inventoryService) CreateInventory(inventory *models.Inventory) (*models.Inventory, error) {
	return service.inventoryRepository.Insert(inventory)
}

func (service *inventoryService) ListAllInventory() ([]models.Inventory, error) {
	return service.inventoryRepository.FindAll()
}

func (service *inventoryService) DeleteInventory(inventoryID primitive.ObjectID) error {
	return service.inventoryRepository.Delete(inventoryID)
}

func (service *inventoryService) UpdateInventory(inventory *models.Inventory) error {
	return service.inventoryRepository.Update(inventory)
}
