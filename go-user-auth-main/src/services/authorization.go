package services

import (
	"go-user-auth/config"
	"go-user-auth/db/repository"
	model "go-user-auth/models"
)

type AuthorizationService interface {
	AuthenticateToken(clientId, clientSecret, code string) (bool, error)
	RegisterClient(clientId, clientSecret string) (bool, error)
	GetClientByID(clientId string) (*model.Client, error)
	SaveAuthorizationCode(authCode model.AuthorizationCode) error
}

type authorizationService struct {
	config   config.AccountsConfig
	authRepo repository.AuthorizationRepository
}

func NewAuthorizationService(config config.AccountsConfig, authRepo repository.AuthorizationRepository) AuthorizationService {
	return &authorizationService{
		config:   config,
		authRepo: authRepo,
	}
}

func (auth *authorizationService) AuthenticateToken(clientId, clientSecret, code string) (bool, error) {
	return auth.authRepo.AuthenticateToken(clientId, clientSecret, code)
}

func (auth *authorizationService) RegisterClient(clientId, clientSecret string) (bool, error) {
	return auth.authRepo.RegisterClient(clientId, clientSecret)
}

func (auth *authorizationService) GetClientByID(clientId string) (*model.Client, error) {
	return auth.authRepo.GetClientByID(clientId)
}

func (auth *authorizationService) SaveAuthorizationCode(authCode model.AuthorizationCode) error {
	return auth.authRepo.SaveAuthorizationCode(authCode)
}
