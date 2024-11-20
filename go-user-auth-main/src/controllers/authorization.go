package controllers

import (
	"encoding/json"
	"go-user-auth/config"
	model "go-user-auth/models"
	"go-user-auth/responses"
	"go-user-auth/services"
	"go-user-auth/utils"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type AuthorizationController interface {
	Token(c *fiber.Ctx) error
	ValidateToken(c *fiber.Ctx) error

	RegisterClient(c *fiber.Ctx) error
	AuthorizeClient(c *fiber.Ctx) error
	Callback(c *fiber.Ctx) error
	ExchangeToken(c *fiber.Ctx) error
}

type authorizationController struct {
	config  config.AccountsConfig
	authLib services.AuthorizationService
}

func NewAuthenticationHandler(config config.AccountsConfig, authLib services.AuthorizationService) AuthorizationController {
	return &authorizationController{
		config:  config,
		authLib: authLib,
	}
}

type Claims struct {
	ClientId string `json:"client_id"`
	jwt.StandardClaims
}

// Create token
// @Summary Create token
// @Description Create token
// @Tags Authentication
// @Produce json
// @Accept application/x-www-form-urlencoded
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} responses.TokenResponse
// @Failure 401 {object} responses.ErrorResponse "Invalid credentials"
// @Failure 500 {object} responses.ErrorResponse "Failed to create token"
// @Router /oauth/token [post]
func (ctrl *authorizationController) Token(c *fiber.Ctx) error {
	clientId := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	// Check if the client_id and client_secret are empty
	log.Println("client_id: ", clientId)
	if clientId == "" || clientSecret == "" {
		log.Println("client_id and client_secret are required")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "client_id and client_secret are required"})
	}

	// Authenticate the user (you need to implement this function)
	status, err := ctrl.authLib.AuthenticateToken(clientId, clientSecret, "auth-code")
	if !status {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to authenticate user", "error": err.Error()})
	}

	expirationTime := time.Now().Add(config.TokenExpiry)
	claims := &Claims{
		ClientId: clientId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.JWTSecret)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create token"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func (ctrl *authorizationController) ValidateToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	c.Locals("clinet_id", claims.ClientId)
	return c.Next()
}

// RegisterClient godoc
// @Summary Register a new client
// @Description Registers a new client with the authorization server.
// @Tags authorization
// @Accept  json
// @Produce  json
// @Param client body model.RegisterClientRequest true "Client details"
// @Success 201 {object} responses.RegisterClientResponse "Client registered successfully"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Router /register-client [post]
func (authorizationController *authorizationController) RegisterClient(c *fiber.Ctx) error {

	var registerClientRequest model.RegisterClientRequest
	json.Unmarshal(c.Body(), &registerClientRequest)

	clientSecret, err := utils.GenerateRandomString(32)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to generate client secret")
	}

	log.Println("client_id: ", registerClientRequest.ClientID)
	if registerClientRequest.ClientID == "" || clientSecret == "" {
		log.Println("client_id and client_secret are required")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "client_id and client_secret are required"})
	}
	// Store the client credentials
	status, err := authorizationController.authLib.RegisterClient(registerClientRequest.ClientID, clientSecret)
	if !status {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register client"})
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to register client", "error": err.Error()})
	}

	response := &model.RegisterClientResponse{
		ClientID:     registerClientRequest.ClientID,
		ClientSecret: clientSecret,
	}

	return c.Status(http.StatusOK).JSON(responses.RegisterClientResponse{
		Status:  http.StatusOK,
		Message: model.Success,
		Data:    response,
	})
}

// AuthorizeClient godoc
// @Summary Authorize a client
// @Description Authorizes a client to access the user's resources and provides an authorization code.
// @Tags authorization
// @Accept  json
// @Produce  json
// @Param client_id query string true "Client ID"
// @Param redirect_uri query string true "Redirect URI"
// @Param response_type query string true "Response Type"
// @Success 302 {object} responses.AuthorizeClientResponse "Client authorized successfully"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Router /authorize [get]
func (authorizationController *authorizationController) AuthorizeClient(c *fiber.Ctx) error {
	// Extract query parameters
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")
	responseType := c.Query("response_type")

	// Validate the response type (only 'code' is allowed in this flow)
	if responseType != "code" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid_request",
			"message": "Invalid response type, expected 'code'.",
		})
	}

	// Fetch client data from MongoDB to validate the client ID
	client, err := authorizationController.authLib.GetClientByID(clientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "server_error",
			"message": "Unable to retrieve client data",
		})
	}

	// Check if the client exists and if the redirect URI matches
	if client == nil || client.RedirectURI != redirectURI {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid_client",
			"message": "Client ID or Redirect URI is invalid.",
		})
	}

	// Generate a unique authorization code
	authCode, err := utils.GenerateRandomString(20) // Generate a 20-character authorization code
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "server_error",
			"message": "Failed to generate authorization code.",
		})
	}

	// Store the authorization code along with the client ID in MongoDB
	authorizationCode := model.AuthorizationCode{
		Code:      authCode,
		ClientID:  clientID,
		ExpiresAt: time.Now().Add(10 * time.Minute).Unix(), // Set expiration time for 10 minutes
	}

	err = authorizationController.authLib.SaveAuthorizationCode(authorizationCode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "server_error",
			"message": "Failed to store authorization code.",
		})
	}

	// Redirect the user to the redirect URI with the authorization code
	return c.Redirect(redirectURI + "?code=" + authCode)
}

// callback godoc
// @Summary Callback after authorization
// @Description Callback after the user authorizes the client.
// @Tags authorization
// @Accept  json
// @Produce  json
// @Param code query string true "Authorization Code"
// @Success 200 {object} responses.AuthorizeClientResponse "Callback successful"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Router /callback [get]
func (authorizationController *authorizationController) Callback(c *fiber.Ctx) error {
	// Retrieve the authorization code from query parameters
	authCode := c.Query("code")

	// Simulate processing the authorization code
	if authCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Authorization code missing",
		})
	}

	// Return the authorization code for further processing
	return c.Status(http.StatusOK).JSON(responses.AuthorizeClientResponse{
		Message:           "Authorization code received",
		AuthorizationCode: authCode,
	})
}

// ExchangeToken godoc
// @Summary Exchange authorization code for access token
// @Description Exchanges the authorization code for an access token.
// @Tags authorization
// @Accept  json
// @Produce  json
// @Param client_id query string true "Client ID"
// @Param client_secret query string true "Client Secret"
// @Param code query string true "Authorization Code"
// @Success 200 {object} responses.ExchangeTokenResponse "Token exchanged successfully"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Router /token [post]
func (authorizationController *authorizationController) ExchangeToken(c *fiber.Ctx) error {
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")
	code := c.FormValue("code")
	// // Verify client credentials and authorization code
	// if secret, ok := clients[clientID]; !ok || secret != clientSecret || code != "auth-code" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "invalid_client_or_code",
	// 	})
	// }

	status, err := authorizationController.authLib.AuthenticateToken(clientID, clientSecret, code)
	if !status {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to authenticate user"})
	}

	// Generate JWT token as the access token
	token, err := utils.GenerateJWT("user1")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Token generation failed")
	}

	return c.JSON(responses.ExchangeTokenResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   900, // 15 minutes
	})
}
