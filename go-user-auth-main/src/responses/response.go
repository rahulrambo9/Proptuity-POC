package responses

import model "go-user-auth/models"

// UserResponse represents the response structure for user-related operations.
type UserResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Error   error       `json:"error,omitempty"`
	Data    *model.User `json:"data,omitempty"`
}

// Struct to return all users
type UsersResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message,omitempty"`
	Error   error         `json:"error,omitempty"`
	Data    *[]model.User `json:"data,omitempty"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RegisterClientResponse struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message,omitempty"`
	Error   error                         `json:"error,omitempty"`
	Data    *model.RegisterClientResponse `json:"data,omitempty"`
}

type AuthorizeClientResponse struct {
	Message           string `json:"message"`
	AuthorizationCode string `json:"authorization_code"`
}

// c.JSON(fiber.Map{
// 	"access_token": token,
// 	"token_type":   "Bearer",
// 	"expires_in":   900, // 15 minutes
// })

type ExchangeTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type ApplicationResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message,omitempty"`
	Error   error              `json:"error,omitempty"`
	Data    *model.Application `json:"data,omitempty"`
}

// Struct to return all Applications
type ApplicationsResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message,omitempty"`
	Error   error                `json:"error,omitempty"`
	Data    *[]model.Application `json:"data,omitempty"`
}
