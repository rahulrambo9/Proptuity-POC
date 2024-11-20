package repository

import (
	"context"
	"fmt"
	"go-user-auth/config"
	model "go-user-auth/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApplicationRepository interface {
	RegisterApplication(ctx context.Context, application *model.Application) (*model.Application, error)
	GetApplicationByClientID(ctx context.Context, clientId string) (*model.Application, error)
	GetApplicationByTag(ctx context.Context, tag string) (*model.Application, error)
	GetAllApplications(ctx context.Context) (*[]model.Application, error)
	DeleteApplication(ctx context.Context, clientId string) (bool, error)
}

type applicationRepository struct {
	cfg             config.AccountsConfig
	mongoRepository MongoRepository
}

func NewApplicationRepository(config config.AccountsConfig, mongoRepository MongoRepository) ApplicationRepository {
	return &applicationRepository{
		cfg:             config,
		mongoRepository: mongoRepository,
	}
}

func (r *applicationRepository) RegisterApplication(ctx context.Context, application *model.Application) (*model.Application, error) {
	log.Println("Register Application in repository")
	application.Id = primitive.NewObjectID()
	result, err := r.mongoRepository.InsertOne(ctx, r.cfg.MongoApplicationCollection, application)
	if err != nil {
		log.Printf("Error while registering application : %+v \n", err)
		return nil, err
	}
	log.Println("User added successfully: ", fmt.Sprintf("%+v", result.InsertedID))
	return application, nil
}

func (r *applicationRepository) GetApplicationByClientID(ctx context.Context, clientId string) (*model.Application, error) {
	log.Println("Get application by client id in repository")
	var application model.Application

	err := r.mongoRepository.FindOne(ctx, r.cfg.MongoApplicationCollection, bson.M{"clientId": clientId}, &application)
	if err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *applicationRepository) GetApplicationByTag(ctx context.Context, tag string) (*model.Application, error) {
	log.Println("Get application by tag in repository")
	var application model.Application

	err := r.mongoRepository.FindOne(ctx, r.cfg.MongoApplicationCollection, bson.M{"tag": tag}, &application)
	if err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *applicationRepository) GetAllApplications(ctx context.Context) (*[]model.Application, error) {
	log.Println("Get all applications in repository")
	var applications []model.Application

	err := r.mongoRepository.FindMany(ctx, r.cfg.MongoApplicationCollection, bson.M{}, &applications)
	if err != nil {
		log.Printf("Error while fetching all applications: %+v \n", err)
		return nil, err
	}

	return &applications, nil
}

func (r *applicationRepository) DeleteApplication(ctx context.Context, clientId string) (bool, error) {
	log.Println("Delete application in repository")

	_, err := r.mongoRepository.DeleteOne(ctx, r.cfg.MongoApplicationCollection, bson.M{"clientId": clientId})
	if err != nil {
		log.Printf("Error while deleting application: %+v \n", err)
		return false, err
	}

	return true, nil
}
