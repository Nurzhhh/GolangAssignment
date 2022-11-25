package services

import (
	"GolangProject/dto"
	"GolangProject/models"
	"GolangProject/repositories"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type OrderService interface {
	Create(b dto.OrderCreateDTO) models.Order
	Update(b dto.OrderUpdateDTO) models.Order
	Delete(b models.Order)
	All() []models.Order
	Show(orderID uint64) models.Order
	IsAllowedToEdit(userID string, orderID uint64) bool
}

type orderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepo,
	}
}

func (service *orderService) Create(o dto.OrderCreateDTO) models.Order {
	order := models.Order{}
	err := smapping.FillStruct(&order, smapping.MapFields(&o))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.orderRepository.CreateOrder(order)
	return res
}

func (service *orderService) Update(b dto.OrderUpdateDTO) models.Order {
	order := models.Order{}
	err := smapping.FillStruct(&order, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.orderRepository.UpdateOrder(order)
	return res
}

func (service *orderService) Delete(o models.Order) {
	service.orderRepository.DeleteOrder(o)
}

func (service *orderService) All() []models.Order {
	return service.orderRepository.AllOrder()
}

func (service *orderService) Show(orderID uint64) models.Order {
	return service.orderRepository.Show(orderID)
}

func (service *orderService) IsAllowedToEdit(userID string, orderID uint64) bool {
	b := service.orderRepository.Show(orderID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
