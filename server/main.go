package main

import (
	"os"
	"server/repository"
	"server/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect DB and intiliaze repositories
	dbconnection := repository.ConnectDB()
	inventoryRepository := repository.NewInventoryRepository(dbconnection)
	shipmentRepository := repository.NewShipmentRepository(dbconnection)

	// Init services
	locationService := service.NewLocationService()
	inventoryService := service.NewInventoryService(inventoryRepository, locationService)
	shipmentService := service.NewShipmentService(shipmentRepository, inventoryRepository, locationService)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	// Inventory Management API
	router.GET("/inventory", inventoryService.GetInventory)
	router.GET("/inventory/:id", inventoryService.GetInventoryById)
	router.POST("/inventory", inventoryService.AddInventory)
	router.DELETE("/inventory/:id", inventoryService.DeleteInventory)
	router.PUT("/inventory/:id", inventoryService.UpdateInventory)

	// Shipment Management API
	router.GET("/shipments", shipmentService.GetShipments)
	router.GET("/shipment/:id", shipmentService.GetShipmentById)
	router.POST("/shipment", shipmentService.CreateShipment)
	router.POST("/shipment/:id/deliver", shipmentService.DeliverShipment)
	router.POST("/shipment/:id/receive", shipmentService.ReceiveShipment)

	// Location Info API
	router.GET("/cities", locationService.GetCities)
	router.GET("/locations", locationService.GetLocationInfo)

	router.Run(":" + port)
}
