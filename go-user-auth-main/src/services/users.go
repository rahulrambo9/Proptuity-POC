package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go-user-auth/config"
	"go-user-auth/db/repository"
	model "go-user-auth/models"
	"go-user-auth/utils"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type UserService interface {
	Get(ctx context.Context, email string) (*model.User, error)
	Add(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, userId int, user *model.User) (*model.User, error)
	Delete(ctx context.Context, userId int) (bool, error)
	GetAll(ctx context.Context) (*[]model.User, error)
	Login(ctx context.Context, email, password string) (model.User, bool, error)
	ResetPassword(email, password string) (bool, error)
	SaveInviteParams(ctx context.Context, inviteParams model.InviteParams) error
	SendEmail(ctx context.Context, baseUrl, magicUrl, email, firstName, LastName string) error
	CallApi(ctx context.Context, url, method string, headers map[string]string, body []byte) ([]byte, error)
	GetInviteParams(ctx context.Context, Uuid string) (model.InviteParams, error)
}

type userService struct {
	config   config.AccountsConfig
	userRepo repository.UserRepository
}

func NewUserService(config config.AccountsConfig, userRepo repository.UserRepository) UserService {
	return &userService{
		config:   config,
		userRepo: userRepo,
	}
}

func (us *userService) Get(ctx context.Context, email string) (*model.User, error) {
	return us.userRepo.GetUser(ctx, email)
}

func (us *userService) Add(ctx context.Context, user *model.User) (*model.User, error) {
	log.Println("Add user in service")
	// Check if the user already exists
	// existingUser, _ := us.Get(ctx, user.UserId)

	filter := bson.M{"email": user.Email}
	projection := bson.M{}
	existingUser, _ := us.userRepo.GetUsersGeneric(ctx, filter, projection)

	if existingUser != nil {
		return &existingUser[0], errors.New("User already exists")
	}
	user.Password = utils.Encrypt(user.Password)
	user.CreatedTS = utils.GetEpochTime()
	return us.userRepo.CreateUser(ctx, user)
}

func (us *userService) Update(ctx context.Context, userId int, user *model.User) (*model.User, error) {
	user.UpdatedTS = utils.GetEpochTime()
	return us.userRepo.UpdateUser(ctx, userId, user)
}

func (us *userService) Delete(ctx context.Context, userId int) (bool, error) {
	return us.userRepo.DeleteUser(ctx, userId)
}

func (us *userService) GetAll(ctx context.Context) (*[]model.User, error) {
	return us.userRepo.GetAllUsers(ctx)
}

func (auth *userService) Login(ctx context.Context, email, password string) (model.User, bool, error) {
	user, status, err := auth.userRepo.Login(ctx, email, password)
	if err != nil {
		log.Printf("Error while logging in: %+v \n", err)
		return model.User{}, false, err
	}

	return user, status, nil
}

func (auth *userService) ResetPassword(email, password string) (bool, error) {
	return auth.userRepo.ResetPassword(email, password)
}

func (auth *userService) SaveInviteParams(ctx context.Context, inviteParams model.InviteParams) error {
	return auth.userRepo.SaveInviteParams(ctx, inviteParams)
}

func (auth *userService) SendEmail(ctx context.Context, baseUrl, magicUrl, email, firstName, LastName string) error {
	// we need to call the api on the baseUrl
	response, err := auth.CallApi(ctx, baseUrl, "POST", map[string]string{}, []byte{})
	if err != nil {
		log.Printf("Error while sending email: %+v \n", err)
		return err
	}
	log.Println("Response from email service: ", string(response))
	return nil
}

func (auth *userService) CallApi(ctx context.Context, url, method string, headers map[string]string, body []byte) ([]byte, error) {
	// Create a new HTTP request with the provided context
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set headers for the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Initialize the HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for the request
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	// Read the response body
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (auth *userService) GetInviteParams(ctx context.Context, Uuid string) (model.InviteParams, error) {
	return auth.userRepo.GetInviteParams(ctx, Uuid)
}
