package services

import (
	"context"
	"go-user-auth/config"
)

type HealthService interface {
	GetHealth(ctx context.Context) (string, error)
}

type healthService struct {
	config config.AccountsConfig
}

func NewHealthService(config config.AccountsConfig) HealthService {

	return &healthService{
		config: config,
	}
}

func (healthSvc *healthService) GetHealth(ctx context.Context) (string, error) {
	// services := make(map[string]string)
	// Todo: Fetch other service health
	// services["USERAPI"] = userapi.GetAppHealthUrl(healthSvc.config.AppEnv)

	resp := `{"status": "UP"}`
	return resp, nil
}
