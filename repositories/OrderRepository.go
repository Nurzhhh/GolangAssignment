package repositories

import (
	"GolangProject/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(b models.Order) models.Order
	UpdateOrder(b models.Order) models.Order
	DeleteOrder(b models.Order)
	AllOrder() []models.Order
	Show(productID uint64) models.Order
}

type orderConnection struct {
	connection *gorm.DB
}

func NewOrderRepository(dbConn *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: dbConn,
	}
}

func (db *orderConnection) CreateOrder(o models.Order) models.Order {
	db.connection.Save(&o)
	db.connection.Preload("User").Find(&o)
	return o
}

func (db *orderConnection) UpdateOrder(o models.Order) models.Order {
	db.connection.Save(&o)
	db.connection.Find(&o)
	return o
}

func (db *orderConnection) DeleteOrder(o models.Order) {
	db.connection.Delete(&o)
}

func (db *orderConnection) Show(orderID uint64) models.Order {
	var order models.Order
	db.connection.Find(&order, orderID)
	return order
}

func (db *orderConnection) AllOrder() []models.Order {
	var orders []models.Order
	db.connection.Preload("User").Find(&orders)
	return orders
}
