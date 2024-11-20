# Go-User-Auth Service

This repository contains the Go-User-Auth service, which is responsible for handling user management, authorization, and authentication using OAuth2. It includes various endpoints for client authorization, token management, and user management functionalities.

## Table of Contents

- [Go-User-Auth Service](#go-user-auth-service)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Service](#running-the-service)
  - [API Documentation](#api-documentation)
    - [Authorization APIs](#authorization-apis)
    - [Order APIs](#order-apis)
    - [Property APIs](#property-apis)
    - [User Management APIs](#user-management-apis)
  - [Environment Variables](#environment-variables)
  - [License](#license)

## Prerequisites

Before running this project, ensure you have the following installed:

- Go (>=1.19)
- MongoDB
- Kafka (for event streaming)
- Redis
- Swagger for API documentation

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-user-auth.git
   cd go-user-auth

2. Install dependencies: 
    ```bash
    make install

3. Run the tests to ensure everything is set up correctly:
    ```bash
    make test

4. Format and lint the code:
    ```bash
    make fmt
    make vet

## Running the Service

1. To run with Docker:

    - Build the Docker image:
        ```bash
        docker build -t go-user-auth .

    - Run the container:
        ```bash
        docker run -p 8080:8080 go-user-auth

## API Documentation
The Go-User-Auth service exposes various APIs for authorization, user management, property, and order handling.

### Authorization APIs
    
a. Authorize Client: `GET /authorize`

This API authorizes a client to access user resources and returns an authorization code.

Parameters:

- `client_id` (query) - Required
- `redirect_uri` (query) - Required
- `response_type` (query) - Required


b. Callback: `GET /callback`

This API handles the callback after user authorization.

Parameters:

- `code` (query) - Required

c. Exchange Token: `POST /token`

This API exchanges the authorization code for an access token.

Parameters:

- `client_id` (query) - Required
- `client_secret` (query) - Required
- `code` (query) - Required

### User Management APIs

a. Login: POST `/user/login`

This API allows a user to log in using their credentials.

Example request body:
    
```json
{
"username": "john_doe",
"password": "password123"
}
```


b. Create User: POST `/v1/users`

This API creates a new user in the system.

Example request body:

```json
{
  "email": "john.doe@example.com",
  "location": "New York",
  "name": "John Doe",
  "password": "password123",
  "phone": "1234567890",
  "username": "john_doe"
}
```

c. Update User: PUT `/v1/users/{userid}`

This API updates an existing user by their ID.

Path parameters:
```bash
userid (integer) - Required
```

Example request body:

```json
{
  "email": "john.doe@example.com",
  "location": "New York",
  "name": "John Doe",
  "password": "newpassword",
  "phone": "1234567890",
  "username": "john_doe"
}
```

d. Delete User: DELETE `/v1/users/{userid}`

This API deletes a user by their ID.

Path parameters:
```bash
userid (integer) - Required
```    

e. Login API: POST `/api/v1/auth/login`

Example request Body
```json
{
  "username": "string",
  "password": "string"
}
```

Example response Body
```json
{
  "access_token": "string",
  "refresh_token": "string",
  "expires_in": 3600
}
```

f. Password Reset API: POST `/api/v1/auth/password-reset`

Example request Body
```json
{
  "email": "string"
}
```

Example response Body
```json
{
  "message": "Password reset link sent to email."
}
```

