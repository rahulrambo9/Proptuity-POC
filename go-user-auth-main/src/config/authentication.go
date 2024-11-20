package config

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	KeyId       = os.Getenv("KEY_ID")
	JWTSecret   = []byte(os.Getenv("KEY_SECRET"))
	TokenExpiry = time.Hour * 24
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
