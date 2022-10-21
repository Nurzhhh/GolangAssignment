package services

import (
	"GolangProject/dto"
	"GolangProject/models"
	"GolangProject/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type ProductService interface {
	Insert(b dto.ProductCreateDTO) models.Product
	Update(b dto.ProductUpdateDTO) models.Product
	Delete(b models.Product)
	All() []models.Product
	Show(productID uint64) models.Product
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (service *productService) Insert(b dto.ProductCreateDTO) models.Product {
	product := models.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.InsertProduct(product)
	return res
}

func (service *productService) Update(b dto.ProductUpdateDTO) models.Product {
	product := models.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.UpdateProduct(product)
	return res
}

func (service *productService) Delete(b models.Product) {
	service.productRepository.DeleteProduct(b)
}

func (service *productService) All() []models.Product {
	return service.productRepository.AllProduct()
}

func (service *productService) Show(productId uint64) models.Product {
	return service.productRepository.Show(productId)
}
