package services

import (
	"context"
	"go-user-auth/config"
)

type OrderService interface {
	CreateOrderService(ctx context.Context) (string, error)
	GetOrderService(ctx context.Context) (string, error)
}

type orderService struct {
	config config.AccountsConfig
}

func NewOrderService(config config.AccountsConfig) OrderService {
	return &orderService{
		config: config,
	}
}

func (orderService *orderService) CreateOrderService(ctx context.Context) (string, error) {
	// TODO: Logic to create order
	return "Order created successfully", nil
}

func (orderService *orderService) GetOrderService(ctx context.Context) (string, error) {
	// TODO: Logic to get order
	return "Order details retrieved successfully", nil
}
