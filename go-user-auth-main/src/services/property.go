package services

import (
	"context"
	"go-user-auth/config"
)

type PropertyService interface {
	CreatePropertyService(ctx context.Context) (string, error)
	GetPropertiesService(ctx context.Context) (string, error)
}

type propertyService struct {
	config config.AccountsConfig
}

func NewPropertyService(config config.AccountsConfig) PropertyService {
	return &propertyService{
		config: config,
	}
}

func (propertyService *propertyService) CreatePropertyService(ctx context.Context) (string, error) {
	// Logic to create property
	return "Property created successfully", nil
}

func (propertyService *propertyService) GetPropertiesService(ctx context.Context) (string, error) {
	// Logic to get properties
	return "Properties retrieved successfully", nil
}
