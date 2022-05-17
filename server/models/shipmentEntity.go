package models

type Shipment struct {
	id string
	label string
	inventoryId string
	quantity map[string]int
	status string
}

type ShipmentService interface {
	CreateShipment(shipment *Shipment) (*Shipment, error)
	ListAllShipment() ([]Shipment, error)
	RecieveShipment(shipment *Shipment) error
}

type ShipmentRepository interface {
	Insert(shipment *Shipment) (*Shipment, error)
	FindOne(id string) (*Shipment, error)
	FindAll() ([]Shipment, error)
	Delete(shipmentID string) error
	UpdateStock(inventoryId string, quantity map[string]int) error
	UpdateStatus(status string, shipmentID string) error
}