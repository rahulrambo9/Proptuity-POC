package controllers

import (
	"go-user-auth/config"
	"go-user-auth/services"

	"github.com/gofiber/fiber/v2"
)

type OrderController interface {
	CreateOrder(c *fiber.Ctx) error
	GetOrder(c *fiber.Ctx) error
}

type orderController struct {
	config   config.AccountsConfig
	orderSvc services.OrderService
}

func NewOrderController(config config.AccountsConfig, orderSvc services.OrderService) OrderController {
	return &orderController{
		config:   config,
		orderSvc: orderSvc,
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Creates a new order for a property purchase or rental.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body model.CreateOrderRequest true "Order details"
// @Success 201 {object} model.OrderResponse "Order created successfully"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Router /orders [post]
func (orderController *orderController) CreateOrder(c *fiber.Ctx) error {
	return c.SendString("Create Order")
}

// GetOrder godoc
// @Summary Get details of an order
// @Description Retrieves the details of an order by its ID.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Success 200 {object} model.OrderResponse "Order details retrieved successfully"
// @Failure 404 {object} model.ErrorResponse "Order not found"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Router /orders/{id} [get]
func (orderController *orderController) GetOrder(c *fiber.Ctx) error {
	return c.SendString("Get Order")
}
