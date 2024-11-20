package model

type Property struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Price    int    `json:"price"`
	Location string `json:"location"`
}

type Order struct {
	ID         string `json:"id"`
	PropertyID string `json:"property_id"`
	BuyerID    string `json:"buyer_id"`
	Price      int    `json:"price"`
	Status     string `json:"status"`
}

// CreateOrderRequest represents the body of the order creation request
type CreateOrderRequest struct {
	PropertyID string `json:"property_id" example:"123"`
	BuyerID    string `json:"buyer_id" example:"456"`
	OfferPrice int    `json:"offer_price" example:"1000000"`
	OrderType  string `json:"order_type" example:"buy"` // or "rent"
}

// OrderResponse represents the response of the order creation or retrieval
type OrderResponse struct {
	ID         string `json:"id" example:"789"`
	PropertyID string `json:"property_id" example:"123"`
	BuyerID    string `json:"buyer_id" example:"456"`
	OfferPrice int    `json:"offer_price" example:"1000000"`
	Status     string `json:"status" example:"pending"`
}

// CreatePropertyRequest represents the body of the property creation request
type CreatePropertyRequest struct {
	Name        string `json:"name" example:"Ocean View Apartment"`
	Type        string `json:"type" example:"apartment"`
	Price       int    `json:"price" example:"3000000"`
	Location    string `json:"location" example:"New York"`
	Bedrooms    int    `json:"bedrooms" example:"3"`
	Bathrooms   int    `json:"bathrooms" example:"2"`
	SquareFeet  int    `json:"square_feet" example:"1200"`
	Description string `json:"description" example:"A spacious apartment with ocean view"`
}

// PropertyResponse represents the response for property-related requests
type PropertyResponse struct {
	ID          string `json:"id" example:"property123"`
	Name        string `json:"name" example:"Ocean View Apartment"`
	Type        string `json:"type" example:"apartment"`
	Price       int    `json:"price" example:"3000000"`
	Location    string `json:"location" example:"New York"`
	Bedrooms    int    `json:"bedrooms" example:"3"`
	Bathrooms   int    `json:"bathrooms" example:"2"`
	SquareFeet  int    `json:"square_feet" example:"1200"`
	Description string `json:"description" example:"A spacious apartment with ocean view"`
}

// ErrorResponse represents a generic error response
type ErrorResponse struct {
	Message string `json:"message" example:"Error occurred"`
}

type RegisterClientRequest struct {
	ClientID     string `json:"client_id" example:"client123"`
	ClientSecret string `json:"client_secret" example:" secret123"`
}

type RegisterClientResponse struct {
	ClientID     string `json:"client_id" example:"client123"`
	ClientSecret string `json:"client_secret" example:" secret123"`
}
