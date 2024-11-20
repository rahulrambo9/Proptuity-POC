package services

import (
	"context"
	"go-user-auth/config"
	"go-user-auth/db/repository"
	model "go-user-auth/models"
)

type ApplicationService interface {
	RegisterApplication(ctx context.Context, application *model.Application) (*model.Application, error)
	GetApplicationByClientID(ctx context.Context, clientId string) (*model.Application, error)
	GetApplicationByTag(ctx context.Context, tag string) (*model.Application, error)
	GetAllApplications(ctx context.Context) (*[]model.Application, error)
	DeleteApplication(ctx context.Context, clientId string) (bool, error)
}

type applicationService struct {
	config          config.AccountsConfig
	applicationRepo repository.ApplicationRepository
}

func NewApplicationService(config config.AccountsConfig, applicationRepo repository.ApplicationRepository) ApplicationService {
	return &applicationService{
		config:          config,
		applicationRepo: applicationRepo,
	}
}

func (s *applicationService) RegisterApplication(ctx context.Context, application *model.Application) (*model.Application, error) {
	return s.applicationRepo.RegisterApplication(ctx, application)
}

func (s *applicationService) GetApplicationByClientID(ctx context.Context, clientId string) (*model.Application, error) {
	return s.applicationRepo.GetApplicationByClientID(ctx, clientId)
}

func (s *applicationService) GetApplicationByTag(ctx context.Context, tag string) (*model.Application, error) {

	return s.applicationRepo.GetApplicationByTag(ctx, tag)
}

func (s *applicationService) GetAllApplications(ctx context.Context) (*[]model.Application, error) {
	return s.applicationRepo.GetAllApplications(ctx)
}

func (s *applicationService) DeleteApplication(ctx context.Context, clientId string) (bool, error) {
	return s.applicationRepo.DeleteApplication(ctx, clientId)
}
