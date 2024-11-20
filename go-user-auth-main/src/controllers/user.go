package controllers

import (
	"encoding/json"
	"fmt"
	"go-user-auth/config"
	"go-user-auth/responses"
	"go-user-auth/stderrors"
	"go-user-auth/utils"
	"log"
	"net/http"
	"net/url"
	"strconv"

	model "go-user-auth/models"
	"go-user-auth/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UserController is the interface for the user controller
type UserController interface {
	Get(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error

	Login(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	VerifyToken(c *fiber.Ctx) error
	Invite(c *fiber.Ctx) error
}

type userController struct {
	config  config.AccountsConfig
	userLib services.UserService
}

func NewUserHandler(config config.AccountsConfig, userLib services.UserService) UserController {
	return &userController{
		config:  config,
		userLib: userLib,
	}
}

// Get retrieves a user by ID.
// @Summary Get a user by ID
// @Description Retrieves a user by the given ID.
// @Tags Users
// @Param userid path int true "User ID"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {object} responses.UserResponse
// @Router /v1/users/{userid} [get]
// @Security ApiKeyAuth
func (userController *userController) Get(c *fiber.Ctx) error {
	// Get the user id from the param
	log.Println("Get user in controller")
	email := c.Params("email", "")

	user, err := userController.userLib.Get(c.UserContext(), email)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   stderrors.ErrNotFound(c.UserContext(), "User not found"),
		})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: model.Success,
		Data:    user,
	})
}

// Add creates a new user.
// @Summary Create a new user
// @Description Creates a new user with the provided details.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 404 {object} responses.UserResponse
// @Router /v1/users [post]
// @Security ApiKeyAuth
func (userController *userController) Add(c *fiber.Ctx) error {
	log.Println("Add user in controller")
	// Unmarshal the request body into the User struct
	var user model.User
	json.Unmarshal(c.Body(), &user)

	response, err := userController.userLib.Add(c.UserContext(), &user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}
	// Respond with the updated list of users
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: model.Success,
		Data:    response,
	})
}

// Update updates an existing user.
// @Summary Update an existing user
// @Description Updates an existing user with the provided details.
// @Tags Users
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Param user body model.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {object} responses.UserResponse
// @Router /v1/users/{userid} [put]
// @Security ApiKeyAuth
func (userController *userController) Update(c *fiber.Ctx) error {
	userId := c.Params("userid", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return c.SendStatus(400)
	}

	// Read the request body
	var updatedUser model.User
	// Unmarshal the request body into the User struct
	err = json.Unmarshal(c.Body(), &updatedUser)
	if err != nil {
		// If the request body is not valid, respond with an error
		return err
	}

	response, err := userController.userLib.Update(c.UserContext(), userIdInt, &updatedUser)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: model.Success,
		Data:    response,
	})
}

// Delete deletes a user by ID.
// @Summary Delete a user by ID
// @Description Deletes a user by the given ID.
// @Tags Users
// @Param userid path int true "User ID"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {object} responses.UserResponse
// @Router /v1/users/{userid} [delete]
// @Security ApiKeyAuth
func (userController *userController) Delete(c *fiber.Ctx) error {
	userId := c.Params("userid", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return c.SendStatus(400)
	}

	status, err := userController.userLib.Delete(c.UserContext(), userIdInt)
	if err != nil {
		return c.SendStatus(404)
	}

	if status {
		return c.Status(http.StatusOK).JSON(responses.UserResponse{
			Status:  http.StatusOK,
			Message: model.Success,
		})
	}

	// If the user is not found, respond with an error
	return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
		Status:  http.StatusNotFound,
		Message: model.Error,
		Error:   stderrors.ErrNotFound(c.UserContext(), "User not found"),
	})
}

// GetAll retrieves all users.
// @Summary Get all users
// @Description Retrieves all users.
// @Tags Users
// @Success 200 {object} responses.UsersResponse
// @Failure 404 {object} responses.UsersResponse
// @Router /v1/users [get]
// @Security ApiKeyAuth
func (userController *userController) GetAll(c *fiber.Ctx) error {
	log.Println("Get all user in controller")
	users, err := userController.userLib.GetAll(c.UserContext())
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}

	return c.Status(http.StatusOK).JSON(responses.UsersResponse{
		Status:  http.StatusOK,
		Message: model.Success,
		Data:    users,
	})
}

// ResetPassword
// @Summary ResetPassword
// @Description ResetPassword
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorResponse "Failed to reset password"
// @Router /user/reset [post]
func (userController *userController) ResetPassword(c *fiber.Ctx) error {
	var user model.User
	json.Unmarshal(c.Body(), &user)
	status, err := userController.userLib.ResetPassword(user.Email, user.Password)

	if err != nil || !status {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to reset password"})
	}

	return c.JSON(fiber.Map{"message": "Password reset successful."})
}

// Login
// @Summary Login
// @Description Login
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorResponse "Bad Request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Router /user/login [post]
func (userController *userController) Login(c *fiber.Ctx) error {
	var loginReq model.LoginRequest
	if err := json.Unmarshal(c.Body(), &loginReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}
	if loginReq.RedirectURI == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "RedirectURI is required"})
	}

	userObj, status, err := userController.userLib.Login(c.UserContext(), loginReq.Email, loginReq.Password)
	if err != nil {
		log.Printf("Error during login: %+v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to login"})
	}
	clientID := c.Get("Client-id")
	if clientID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Client-id header is required"})
	}

	if status {
		accessToken, idToken, refreshToken, err := userController.generateTokens(loginReq.Email, clientID, userObj)
		if err != nil {
			log.Printf("Error generating tokens: %+v\n", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating tokens"})
		}

		redirectURI, err := url.Parse(loginReq.RedirectURI)
		if err != nil {
			log.Printf("Error parsing redirect URI: %+v\n", err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid redirect URI"})
		}

		queryParams := redirectURI.Query()
		queryParams.Set("access_token", accessToken)
		queryParams.Set("id_token", idToken)
		queryParams.Set("refresh_token", refreshToken)
		redirectURI.RawQuery = queryParams.Encode()

		log.Println("Redirecting to:", redirectURI.String())
		return c.Status(http.StatusFound).Redirect(redirectURI.String())
	}
	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}

// SignUp
// @Summary SignUp
// @Description SignUp
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorResponse "Bad Request"
// @Router /signup [post]
func (userController *userController) SignUp(c *fiber.Ctx) error {
	var signUpRequest model.SignUpRequest
	if err := json.Unmarshal(c.Body(), &signUpRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}
	if signUpRequest.RedirectURI == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "RedirectURI is required"})
	}

	var inviteObj model.InviteParams
	inviteUuid := uuid.New()
	inviteObj.Uuid = inviteUuid
	inviteObj.Status = "pending"
	inviteObj.Email = signUpRequest.Email
	inviteObj.CreatedTS = utils.GetEpochTime()
	inviteObj.UpdatedTS = utils.GetEpochTime()

	err := userController.userLib.SaveInviteParams(c.UserContext(), inviteObj)
	if err != nil {
		log.Printf("Error saving invite params: %+v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to save invite params"})
	}

	// Create magic link
	magicUrl := userController.config.Dns + "/go-user-auth/invite/" + inviteUuid.String()

	baseUrl := userController.config.Dns + "/notifications/send"
	err = userController.userLib.SendEmail(c.UserContext(), baseUrl, magicUrl, signUpRequest.Email, signUpRequest.FirstName, signUpRequest.LastName)
	if err != nil {
		log.Printf("Error sending email: %+v\n", err)
		//return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to send email"})
	}

	var user model.User
	user.Email = signUpRequest.Email
	user.Password = signUpRequest.Password
	user.UserType = signUpRequest.UserType
	user.ProfessionalType = signUpRequest.ProfessionalType
	user.FirstName = signUpRequest.FirstName
	user.LastName = signUpRequest.LastName
	user.Zip = signUpRequest.Zip
	user.Phone = signUpRequest.Phone

	_, err = userController.userLib.Add(c.UserContext(), &user)
	if err != nil {
		log.Printf("Error during sign up: %+v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to sign up : " + err.Error()})
	}

	clientID := c.Get("Client-id")
	if clientID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Client-id header is required"})
	}

	accessToken, idToken, refreshToken, err := userController.generateTokens(signUpRequest.Email, clientID, user)
	if err != nil {
		log.Printf("Error generating tokens: %+v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating tokens"})
	}

	redirectURI, err := url.Parse(signUpRequest.RedirectURI)
	if err != nil {
		log.Printf("Error parsing redirect URI: %+v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid redirect URI"})
	}

	queryParams := redirectURI.Query()
	queryParams.Set("access_token", accessToken)
	queryParams.Set("id_token", idToken)
	queryParams.Set("refresh_token", refreshToken)
	redirectURI.RawQuery = queryParams.Encode()

	log.Println("Redirecting to:", redirectURI.String())
	return c.Status(http.StatusFound).Redirect(redirectURI.String())
}

func (userController *userController) generateTokens(email, clientID string, user model.User) (string, string, string, error) {
	privKey := userController.config.KeySecret

	accessToken, err := utils.GenerateJWTToken(email, privKey, "go-user-auth", clientID, model.User{})
	if err != nil {
		return "", "", "", fmt.Errorf("error generating access token: %w", err)
	}

	idToken, err := utils.GenerateJWTToken(email, privKey, "go-user-auth", clientID, user)
	if err != nil {
		return "", "", "", fmt.Errorf("error generating ID token: %w", err)
	}

	refreshToken, err := utils.GenerateJWTToken(email, privKey, "go-user-auth", clientID, model.User{})
	if err != nil {
		return "", "", "", fmt.Errorf("error generating refresh token: %w", err)
	}

	return accessToken, idToken, refreshToken, nil
}

// VerifyToken
// @Summary Verify JWT Token
// @Description Verifies the authenticity and validity of a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Success 200 {object} responses.SuccessResponse
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 400 {object} responses.ErrorResponse "Invalid or expired token"
// @Router /verify-token [get]
func (userController *userController) VerifyToken(c *fiber.Ctx) error {
	// Get the token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Authorization header missing"})
	}

	// Verify and parse the token
	claims, err := utils.VerifyJWTToken(authHeader, userController.config.KeyId)
	if err != nil {
		log.Printf("Token verification failed: %+v \n", err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Return success with token claims (e.g., user email, client ID, etc.)
	return c.JSON(fiber.Map{
		"message": "Token is valid",
		"claims":  claims,
	})
}

// Invite
// @Summary Invite
// @Description Invite
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param uuid path string true "UUID"
// @Success 200 {object} responses.UserResponse
// @Router /invite [post]
func (userController *userController) Invite(c *fiber.Ctx) error {
	// var inviteObj model.InviteParams
	// json.Unmarshal(c.Body(), &inviteObj)

	// Validate the jwt token
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Authorization header missing"})
	}
	_, err := utils.VerifyJWTToken(authHeader, userController.config.KeyId)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	uuid := c.Params("uuid", "")
	if uuid == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Uuid is required"})
	}

	// Get MongoInviteCollection from db to validate if uuid is present and check its expiry
	inviteParams, err := userController.userLib.GetInviteParams(c.UserContext(), uuid)
	if err != nil {
		log.Printf("Error fetching invite params: %+v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid invite link"})
	}

	// Check if the invite is expired && Check if the invite expiry time is greater than the current time
	if inviteParams.Status == "expired" || inviteParams.ExpiresAt < utils.GetEpochTime() {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invite link has expired"})
	}
	return c.JSON(fiber.Map{"message": "Invite sent successfully"})
}
