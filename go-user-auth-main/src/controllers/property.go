package controllers

import (
	"go-user-auth/config"
	"go-user-auth/services"

	"github.com/gofiber/fiber/v2"
)

type PropertyController interface {
	CreateProperty(c *fiber.Ctx) error
	GetProperties(c *fiber.Ctx) error
}

type propertyController struct {
	config      config.AccountsConfig
	propertySvc services.PropertyService
}

func NewPropertyController(config config.AccountsConfig, propertySvc services.PropertyService) PropertyController {
	return &propertyController{
		config:      config,
		propertySvc: propertySvc,
	}
}

// CreateProperty godoc
// @Summary Create a new property
// @Description Onboards a new property for sale or rent.
// @Tags properties
// @Accept  json
// @Produce  json
// @Param property body model.CreatePropertyRequest true "Property details"
// @Success 201 {object} model.PropertyResponse "Property created successfully"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Router /properties [post]
func (propertyController *propertyController) CreateProperty(c *fiber.Ctx) error {
	return c.SendString("Create Property")
}

// GetProperties godoc
// @Summary Get all properties
// @Description Fetches a list of properties based on filters such as location, price range, and property type.
// @Tags properties
// @Accept  json
// @Produce  json
// @Param city query string false "City filter"
// @Param min_price query int false "Minimum price filter"
// @Param max_price query int false "Maximum price filter"
// @Param property_type query string false "Property type filter (e.g., apartment, house)"
// @Success 200 {array} model.PropertyResponse "Properties retrieved successfully"
// @Failure 404 {object} model.ErrorResponse "No properties found"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Router /properties [get]
func (propertyController *propertyController) GetProperties(c *fiber.Ctx) error {
	return c.SendString("Get Properties")
}
