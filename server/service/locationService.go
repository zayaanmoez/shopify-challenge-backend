package service

import (
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type LocationService struct {
	// inventoryRepository *repository.InventoryRepository
	Locations  []string
	Warehouses map[string][]string
	api_key    string
}

var onceLocation sync.Once

// Create an instance of the Shipment service
func NewLocationService() *LocationService {
	var instance *LocationService
	onceLocation.Do(func() {
		instance = new(LocationService)
		instance.Locations = []string{"London", "New York City", "Paris", "Sydney", "Tokyo"}
		instance.Warehouses = map[string][]string{
			"London":        {"Warehouse-London1", "Warehouse-London2"},
			"New York City": {"Warehouse-NewYorkCity1", "Warehouse-NewYorkCity2"},
			"Paris":         {"Warehouse-Paris1", "Warehouse-Paris2"},
			"Sydney":        {"Warehouse-Sydney1", "Warehouse-Sydney2"},
			"Tokyo":         {"Warehouse-Tokyo1", "Warehouse-Tokyo2"},
		}
		instance.api_key = os.Getenv("API_KEY")
	})

	return instance
}

func (service *LocationService) GetCities(c *gin.Context) {
	c.JSON(http.StatusOK, service.Locations)
}

func (service *LocationService) GetLocationInfo(c *gin.Context) {
	c.JSON(http.StatusOK, service.Warehouses)
}

func (service *LocationService) IsValidCity(city string) bool {
	for _, c := range service.Locations {
		if c == city {
			return true
		}
	}
	return false
}

func (service *LocationService) IsValidWarehouse(city string, warehouse string) bool {
	for _, c := range service.Warehouses[city] {
		if c == warehouse {
			return true
		}
	}
	return false
}

func (service *LocationService) GetWarehouses(city string) []string {
	return service.Warehouses[city]
}
