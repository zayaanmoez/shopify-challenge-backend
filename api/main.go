package main

import (
	"github.com/gin-gonic/gin"
	"shopify-challenge/database"
	"shopify-challenge/controllers"
	"shopify-challenge/service"
	"shopify-challenge/repository"
)

func main() {
	connection := database.ConnectDB()
	inventoryRepository := repository.NewInventoryRepository(connection)
	inventoryService := service.NewInventoryService(inventoryRepository)
	inventoryController := controllers.NewInventoryController(inventoryService)

	router := gin.Default()
	router.GET("/inventory", inventoryController.GetInventory)
	router.POST("/inventory", inventoryController.AddInventory)
	router.DELETE("/inventory", inventoryController.DeleteInventory)
	router.PUT("/inventory", inventoryController.UpdateInventory)

	router.Run(":5050")
}