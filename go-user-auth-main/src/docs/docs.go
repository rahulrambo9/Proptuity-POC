// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/authorize": {
            "get": {
                "description": "Authorizes a client to access the user's resources and provides an authorization code.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "Authorize a client",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "client_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Redirect URI",
                        "name": "redirect_uri",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Response Type",
                        "name": "response_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Client authorized successfully",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthorizeClientResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/callback": {
            "get": {
                "description": "Callback after the user authorizes the client.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "Callback after authorization",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Callback successful",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthorizeClientResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/oauth/token": {
            "post": {
                "description": "Create token",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Create token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TokenResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create token",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/orders": {
            "post": {
                "description": "Creates a new order for a property purchase or rental.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "Order details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Order created successfully",
                        "schema": {
                            "$ref": "#/definitions/model.OrderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Retrieves the details of an order by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get details of an order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order details retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/model.OrderResponse"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/properties": {
            "get": {
                "description": "Fetches a list of properties based on filters such as location, price range, and property type.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "properties"
                ],
                "summary": "Get all properties",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City filter",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Minimum price filter",
                        "name": "min_price",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Maximum price filter",
                        "name": "max_price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Property type filter (e.g., apartment, house)",
                        "name": "property_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Properties retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.PropertyResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "No properties found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Onboards a new property for sale or rent.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "properties"
                ],
                "summary": "Create a new property",
                "parameters": [
                    {
                        "description": "Property details",
                        "name": "property",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatePropertyRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Property created successfully",
                        "schema": {
                            "$ref": "#/definitions/model.PropertyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/register-client": {
            "post": {
                "description": "Registers a new client with the authorization server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "Register a new client",
                "parameters": [
                    {
                        "description": "Client details",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterClientRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Client registered successfully",
                        "schema": {
                            "$ref": "#/definitions/responses.RegisterClientResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/token": {
            "post": {
                "description": "Exchanges the authorization code for an access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "Exchange authorization code for access token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "client_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Client Secret",
                        "name": "client_secret",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization Code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token exchanged successfully",
                        "schema": {
                            "$ref": "#/definitions/responses.ExchangeTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/reset": {
            "post": {
                "description": "ResetPassword",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "ResetPassword",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Failed to reset password",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves all users.",
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UsersResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UsersResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new user with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/{userid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves a user by the given ID.",
                "tags": [
                    "Users"
                ],
                "summary": "Get a user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates an existing user with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update an existing user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes a user by the given ID.",
                "tags": [
                    "Users"
                ],
                "summary": "Delete a user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateOrderRequest": {
            "type": "object",
            "properties": {
                "buyer_id": {
                    "type": "string",
                    "example": "456"
                },
                "offer_price": {
                    "type": "integer",
                    "example": 1000000
                },
                "order_type": {
                    "description": "or \"rent\"",
                    "type": "string",
                    "example": "buy"
                },
                "property_id": {
                    "type": "string",
                    "example": "123"
                }
            }
        },
        "model.CreatePropertyRequest": {
            "type": "object",
            "properties": {
                "bathrooms": {
                    "type": "integer",
                    "example": 2
                },
                "bedrooms": {
                    "type": "integer",
                    "example": 3
                },
                "description": {
                    "type": "string",
                    "example": "A spacious apartment with ocean view"
                },
                "location": {
                    "type": "string",
                    "example": "New York"
                },
                "name": {
                    "type": "string",
                    "example": "Ocean View Apartment"
                },
                "price": {
                    "type": "integer",
                    "example": 3000000
                },
                "square_feet": {
                    "type": "integer",
                    "example": 1200
                },
                "type": {
                    "type": "string",
                    "example": "apartment"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Error occurred"
                }
            }
        },
        "model.OrderResponse": {
            "type": "object",
            "properties": {
                "buyer_id": {
                    "type": "string",
                    "example": "456"
                },
                "id": {
                    "type": "string",
                    "example": "789"
                },
                "offer_price": {
                    "type": "integer",
                    "example": 1000000
                },
                "property_id": {
                    "type": "string",
                    "example": "123"
                },
                "status": {
                    "type": "string",
                    "example": "pending"
                }
            }
        },
        "model.PropertyResponse": {
            "type": "object",
            "properties": {
                "bathrooms": {
                    "type": "integer",
                    "example": 2
                },
                "bedrooms": {
                    "type": "integer",
                    "example": 3
                },
                "description": {
                    "type": "string",
                    "example": "A spacious apartment with ocean view"
                },
                "id": {
                    "type": "string",
                    "example": "property123"
                },
                "location": {
                    "type": "string",
                    "example": "New York"
                },
                "name": {
                    "type": "string",
                    "example": "Ocean View Apartment"
                },
                "price": {
                    "type": "integer",
                    "example": 3000000
                },
                "square_feet": {
                    "type": "integer",
                    "example": 1200
                },
                "type": {
                    "type": "string",
                    "example": "apartment"
                }
            }
        },
        "model.RegisterClientRequest": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string",
                    "example": "client123"
                },
                "client_secret": {
                    "type": "string",
                    "example": " secret123"
                }
            }
        },
        "model.RegisterClientResponse": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string",
                    "example": "client123"
                },
                "client_secret": {
                    "type": "string",
                    "example": " secret123"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "salt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.AuthorizeClientResponse": {
            "type": "object",
            "properties": {
                "authorization_code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "responses.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ExchangeTokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "responses.RegisterClientResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.RegisterClientResponse"
                },
                "error": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.TokenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "responses.UserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.User"
                },
                "error": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.UsersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "error": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "",
	Description:      "JWT Authentication header for all requests",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}