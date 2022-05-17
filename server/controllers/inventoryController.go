package controllers

import (
	"net/http"
	"shopify-challenge/models"
	"github.com/gin-gonic/gin"
)

type inventoryController struct {
	inventoryService models.InventoryService
}

type InventoryController interface {
	AddInventory(c *gin.Context)
	GetInventory(c *gin.Context)
	DeleteInventory(c *gin.Context)
	UpdateInventory(c *gin.Context)
}

func NewInventoryController(service models.InventoryService) InventoryController {
	return &inventoryController {
		inventoryService: service,
	}
}

func (controller *inventoryController) AddInventory(c *gin.Context) {

}

func (controller *inventoryController) GetInventory(c *gin.Context) {
	inventory, err := controller.inventoryService.ListAllInventory()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting inventory"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

func (controller *inventoryController) DeleteInventory(c *gin.Context) {

}

func (controller *inventoryController) UpdateInventory(c *gin.Context) {

}
