package repository

import (
	"context"
	"fmt"
	"go-user-auth/config"
	model "go-user-auth/models"
	"go-user-auth/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetUser(ctx context.Context, email string) (*model.User, error)
	GetUsersGeneric(ctx context.Context, filter bson.M, projection bson.M) ([]model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, userId int, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, userId int) (bool, error)
	GetAllUsers(ctx context.Context) (*[]model.User, error)

	Login(ctx context.Context, email, password string) (model.User, bool, error)
	ResetPassword(email, password string) (bool, error)
	SaveInviteParams(ctx context.Context, inviteParams model.InviteParams) error
	GetInviteParams(ctx context.Context, Uuid string) (model.InviteParams, error)
}

type userRepository struct {
	cfg             config.AccountsConfig
	mongoRepository MongoRepository
}

func NewUserRepository(config config.AccountsConfig, mongoRepository MongoRepository) UserRepository {
	return &userRepository{
		cfg:             config,
		mongoRepository: mongoRepository,
	}
}

func (r *userRepository) GetUser(ctx context.Context, email string) (*model.User, error) {
	log.Println("Get user in repository")
	var user model.User

	err := r.mongoRepository.FindOne(ctx, r.cfg.MongoUserCollection, bson.M{"email": email}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUsersGeneric(ctx context.Context, filter bson.M, projection bson.M) ([]model.User, error) {
	log.Println("Get all users by filter in repository")
	var users []model.User
	opts := options.Find().SetProjection(projection)
	err := r.mongoRepository.FindMany(ctx, r.cfg.MongoUserCollection, filter, &users, opts)
	if err != nil {
		log.Printf("Error while fetching all users: %+v \n", err)
		return nil, err
	}

	return users, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	log.Println("Add user in repository")
	user.Id = primitive.NewObjectID()
	result, err := r.mongoRepository.InsertOne(ctx, r.cfg.MongoUserCollection, user)
	if err != nil {
		log.Printf("Error while adding user: %+v \n", err)
		return nil, err
	}
	log.Println("User added successfully: ", fmt.Sprintf("%+v", result.InsertedID))
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, userId int, user *model.User) (*model.User, error) {
	log.Println("Update user in repository")

	_, err := r.mongoRepository.UpdateOne(ctx, r.cfg.MongoUserCollection, bson.M{"userId": userId}, bson.M{"$set": user})
	if err != nil {
		log.Printf("Error while updating user: %+v \n", err)
		return nil, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userId int) (bool, error) {
	log.Println("Delete user in repository")

	result, err := r.mongoRepository.DeleteOne(ctx, r.cfg.MongoUserCollection, bson.M{"userId": userId})
	if result.DeletedCount == 0 {
		log.Printf("User with id %d not found \n", userId)
		return false, nil
	}
	if err != nil {
		log.Printf("Error while deleting user: %+v \n", err)
		return false, err
	}
	return true, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) (*[]model.User, error) {
	log.Println("Get all users in repository")
	var users []model.User

	err := r.mongoRepository.FindMany(ctx, r.cfg.MongoUserCollection, bson.M{}, &users)
	if err != nil {
		log.Printf("Error while fetching all users: %+v \n", err)
		return nil, err
	}

	return &users, nil
}

func (auth *userRepository) Login(ctx context.Context, email, password string) (model.User, bool, error) {
	log.Println("Login in repository")
	var user model.User
	err := auth.mongoRepository.FindOne(ctx, auth.cfg.MongoUserCollection, bson.M{"email": email}, &user)
	if err != nil {
		log.Printf("Error while fetching user: %+v \n", err)
		return model.User{}, false, err
	}
	return user, utils.Compare(password, user.Password), nil
}

func (auth *userRepository) ResetPassword(email, password string) (bool, error) {
	log.Println("ResetPassword in repository")
	var user model.User

	err := auth.mongoRepository.FindOne(context.Background(), auth.cfg.MongoUserCollection, bson.M{"email": email}, &user)
	if err != nil {
		return false, err
	}

	_, err = auth.mongoRepository.UpdateOne(context.Background(), auth.cfg.MongoUserCollection, bson.M{"email": email}, bson.M{"$set": bson.M{"password": password, "UpdatedTS": time.Now()}})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (auth *userRepository) SaveInviteParams(ctx context.Context, inviteParams model.InviteParams) error {
	log.Println("SaveInviteParams in repository")
	_, err := auth.mongoRepository.InsertOne(ctx, auth.cfg.MongoInviteCollection, inviteParams)
	if err != nil {
		log.Printf("Error while inserting invite params: %+v \n", err)
		return err
	}
	return nil
}

func (auth *userRepository) GetInviteParams(ctx context.Context, Uuid string) (model.InviteParams, error) {
	log.Println("GetInviteParams in repository")
	var inviteParams model.InviteParams
	err := auth.mongoRepository.FindOne(ctx, auth.cfg.MongoInviteCollection, bson.M{"uuid": Uuid}, &inviteParams)
	if err != nil {
		log.Printf("Error while fetching invite params: %+v \n", err)
		return model.InviteParams{}, err
	}
	return inviteParams, nil
}
