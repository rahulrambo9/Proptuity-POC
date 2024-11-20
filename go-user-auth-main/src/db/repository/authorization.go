package repository

import (
	"context"
	"fmt"
	"go-user-auth/config"
	model "go-user-auth/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorizationRepository interface {
	AuthenticateToken(clientId, clientSecret, code string) (bool, error)
	RegisterClient(clientId, clientSecret string) (bool, error)
	GetClientByID(clientId string) (*model.Client, error)
	SaveAuthorizationCode(authCode model.AuthorizationCode) error
}

type authorizationRepository struct {
	cfg             config.AccountsConfig
	storageInstance *mongo.Client
}

func NewAuthorizationRepository(config config.AccountsConfig) AuthorizationRepository {
	return &authorizationRepository{
		cfg:             config,
		storageInstance: StorageInstance,
	}
}

func (auth *authorizationRepository) AuthenticateToken(clientId, clientSecret, code string) (bool, error) {
	// if (auth.cfg.ClientId == clientId) && (auth.cfg.ClientSecret == clientSecret) {
	// 	return true, nil
	// }

	collection := GetCollection(auth.storageInstance, auth.cfg.MongoDBName, auth.cfg.MongoClientCollection)
	var client model.Client
	err := collection.FindOne(context.Background(), bson.M{"client_id": clientId, "client_secret": clientSecret}).Decode(&client)
	if err != nil || code != "auth-code" {
		return false, fmt.Errorf("invalid_client_or_code : %v", err.Error())
	}

	return false, nil
}

func (auth *authorizationRepository) RegisterClient(clientId, clientSecret string) (bool, error) {
	collection := GetCollection(auth.storageInstance, auth.cfg.MongoDBName, auth.cfg.MongoClientCollection)
	_, err := collection.InsertOne(context.Background(), bson.M{"clientId": clientId, "clientSecret": clientSecret})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (auth *authorizationRepository) GetClientByID(clientId string) (*model.Client, error) {
	collection := GetCollection(auth.storageInstance, auth.cfg.MongoDBName, auth.cfg.MongoClientCollection)
	var client model.Client
	err := collection.FindOne(context.Background(), bson.M{"client_id": clientId}).Decode(&client)
	if err != nil {
		return &model.Client{}, err
	}

	return &client, nil
}

func (auth *authorizationRepository) SaveAuthorizationCode(authCode model.AuthorizationCode) error {
	collection := GetCollection(auth.storageInstance, auth.cfg.MongoDBName, auth.cfg.MongoAuthCodeCollection)
	_, err := collection.InsertOne(context.Background(), authCode)
	if err != nil {
		return err
	}

	return nil
}
