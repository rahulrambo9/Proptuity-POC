package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Success = "success"
	Error   = "error"
)

type User struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email            string             `json:"email,omitempty"`
	Password         string             `json:"password,omitempty"`
	UserType         string             `json:"user_type,omitempty"`
	ProfessionalType string             `json:"professional_type,omitempty"`
	FirstName        string             `json:"first_name,omitempty"`
	LastName         string             `json:"last_name,omitempty"`
	Zip              string             `json:"zip,omitempty"`
	Phone            string             `json:"phone,omitempty"`
	CreatedTS        int64              `json:"createdTS,omitempty" bson:"createdTS,omitempty"`
	UpdatedTS        int64              `json:"updatedTS,omitempty" bson:"updatedTS,omitempty"`
	UserRoles        []string           `json:"user_roles,omitempty" bson:"user_roles,omitempty"`
	Permissions      []string           `json:"permissions,omitempty" bson:"permissions,omitempty"`
	Applications     []string           `json:"applications,omitempty" bson:"applications,omitempty"`
}

type Client struct {
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

type AuthorizationCode struct {
	Code        string `json:"code,omitempty"`
	ClientID    string `json:"client_id,omitempty"`
	RedirectURI string `json:"redirect_uri,omitempty"`
	ExpiresAt   int64  `json:"expires_at,omitempty"`
}

type LoginRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	RedirectURI string `json:"redirect_uri"` // Client-defined redirect URI
}

type SignUpRequest struct {
	UserId           int    `json:"userId,omitempty" bson:"userId,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	UserType         string `json:"user_type,omitempty"`
	ProfessionalType string `json:"professional_type,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Zip              string `json:"zip,omitempty"`
	Phone            string `json:"phone,omitempty"`
	RedirectURI      string `json:"redirect_uri"` // Client-defined redirect URI
}

type InviteParams struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Uuid      uuid.UUID          `json:"uuid,omitempty"`
	Status    string             `json:"status,omitempty"`
	Email     string             `json:"email,omitempty"`
	ExpiresAt int64              `json:"expires_at,omitempty"`
	CreatedTS int64              `json:"createdTS,omitempty" bson:"createdTS,omitempty"` // todo: update this to createdTS
	UpdatedTS int64              `json:"updatedTS,omitempty" bson:"updatedTS,omitempty"`
}

type Profile struct {
	Id             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
	Bio            string             `json:"bio,omitempty" bson:"bio,omitempty"`
	Address        string             `json:"address,omitempty" bson:"address,omitempty"`
	City           string             `json:"city,omitempty" bson:"city,omitempty"`
	State          string             `json:"state,omitempty" bson:"state,omitempty"`
	Country        string             `json:"country,omitempty" bson:"country,omitempty"`
	ProfilePicture string             `json:"profile_picture,omitempty" bson:"profile_picture,omitempty"`
	CreatedTS      int64              `json:"createdTS,omitempty" bson:"createdTS,omitempty"`
	UpdatedTS      int64              `json:"updatedTS,omitempty" bson:"updatedTS,omitempty"`
}

type Application struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ClientId     string             `json:"client_id,omitempty" bson:"client_id,omitempty"`
	ClientSecret string             `json:"client_secret,omitempty" bson:"client_secret,omitempty"`
	RedirectURI  string             `json:"redirect_uri,omitempty" bson:"redirect_uri,omitempty"`
	CreatedTS    int64              `json:"createdTS,omitempty" bson:"createdTS,omitempty"`
	UpdatedTS    int64              `json:"updatedTS,omitempty" bson:"updatedTS,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	Issuer       string             `json:"issuer,omitempty" bson:"issuer,omitempty"`
}
