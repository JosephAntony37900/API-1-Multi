package service

import (
	"log"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (service *OrderService) ProcessOrder(estado string, tipo bool) error {
	log.Printf("Procesando orden: Estado = %s, Tipo = %t", estado, tipo)
	return nil 
}