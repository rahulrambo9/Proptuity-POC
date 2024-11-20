package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("secret_key")

// JWT claims structure
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Function to generate a secure random string
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// Function to issue JWT tokens
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func JWTToken(privKey *ecdsa.PrivateKey, userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(privKey)
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWTToken creates a JWT token (Access/ID Token) signed by ECDSA private key
func GenerateJWTToken(username string, privKeyStr string, issuer string, audience string) (string, error) {
	privBytes, err := base64.StdEncoding.DecodeString(privKeyStr)
	if err != nil {
		return "", err
	}
	privKey := secp256k1.PrivKeyFromBytes(privBytes)

	// Define claims
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,                                            // Your OAuth server name
			Subject:   "user",                                            // Can use user ID or unique subject
			Audience:  []string{audience},                                // Audience (e.g., client ID)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // 1 hour expiry
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a token object, specifying the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(privKey.ToECDSA())
}
