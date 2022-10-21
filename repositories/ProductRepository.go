package repositories

import (
	"GolangProject/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(b models.Product) models.Product
	UpdateProduct(b models.Product) models.Product
	DeleteProduct(b models.Product)
	AllProduct() []models.Product
	Show(productID uint64) models.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(dbConn *gorm.DB) ProductRepository {
	return &productConnection{
		connection: dbConn,
	}
}

func (db *productConnection) InsertProduct(b models.Product) models.Product {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *productConnection) UpdateProduct(b models.Product) models.Product {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *productConnection) DeleteProduct(b models.Product) {
	db.connection.Delete(&b)
}

func (db *productConnection) Show(productID uint64) models.Product {
	var product models.Product
	db.connection.Find(&product, productID)
	return product
}

func (db *productConnection) AllProduct() []models.Product {
	var products []models.Product
	db.connection.Find(&products)
	return products
}
