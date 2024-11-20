package controllers

// I need to add the test for genrating the jwt token using the ecdsa private key and verifying the token using ecdsa public key

import (
	"encoding/base64"
	"errors"
	"fmt"
	"go-user-auth/helper"
	"testing"
	"time"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	keySecret, keyId, err := GenerateKeyPair()
	fmt.Println("Key Secret: ", keySecret, "  and     Key Id: ", keyId)
	assert.NoError(t, err)

	accessToken, err := GenerateJWTToken("sachin@gamil.com", keySecret, "go-user-auth", "1234")
	fmt.Println("Access Token: ", accessToken)
	assert.NoError(t, err)

	// Verify the token
	// tokenString := strings.TrimPrefix(accessToken, "Bearer ")

	// Verify and parse the token
	claims, err := VerifyJWTToken(accessToken, keyId)
	assert.Error(t, err)
	fmt.Println(claims)
}

func TestGenerateKeyPair(t *testing.T) {
	privKey1, pubKey1, err := helper.GenerateKeyPair()
	assert.NoError(t, err)
	fmt.Println("Private Key 1: ", privKey1)
	fmt.Println("Public Key 1: ", pubKey1)
	fmt.Printf("Public Key Length: %d\n", len(pubKey1))

	privKey2, pubKey2, err := helper.GenerateKeyPair()
	assert.NoError(t, err)
	fmt.Println("Private Key 2: ", privKey2)
	fmt.Println("Public Key 2: ", pubKey2)
	fmt.Printf("Public Key Length: %d\n", len(pubKey1))

	sharedKey1, err := helper.CreateSharedKey(privKey1, pubKey2)
	assert.NoError(t, err)
	fmt.Println("Shared Key 1: ", sharedKey1)

	sharedKey2, err := helper.CreateSharedKey(privKey2, pubKey1)
	assert.NoError(t, err)
	fmt.Println("Shared Key 2: ", sharedKey2)
	assert.Error(t, err)

	sharedKey1Bytes := base64.StdEncoding.EncodeToString(sharedKey1)
	sharedKey2Bytes := base64.StdEncoding.EncodeToString(sharedKey2)

	fmt.Println("Shared Key 1 : ", (sharedKey1Bytes))
	fmt.Println("Shared Key 2 : ", (sharedKey2Bytes))

}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
// 	privKey, err := secp256k1.GeneratePrivateKey()
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	return privKey.ToECDSA(), &privKey.ToECDSA().PublicKey, nil
// }

// func GenerateJWTToken(username string, privKey *ecdsa.PrivateKey, issuer string, audience string) (string, error) {
// 	// Define claims
// 	claims := CustomClaims{
// 		Username: username,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Issuer:    issuer,
// 			Subject:   "user",
// 			Audience:  []string{audience},
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 		},
// 	}

// 	// Create a token object, specifying the signing method
// 	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
// 	return token.SignedString(privKey)
// }

// func VerifyJWTToken(tokenString string, pubKey *ecdsa.PublicKey) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return pubKey, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		return claims, nil
// 	}

//		return nil, errors.New("invalid token claims")
//	}

func GenerateKeyPair() (string, string, error) {
	privKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return "", "", err
	}

	// Encode private key as base64
	dBytes := privKey.ToECDSA().D.Bytes()
	privKeyStr := base64.StdEncoding.EncodeToString(dBytes)

	// Encode public key, including the 0x04 prefix for uncompressed format
	pubKeyBytes := privKey.PubKey().SerializeUncompressed()
	pubKeyStr := base64.StdEncoding.EncodeToString(pubKeyBytes)

	return privKeyStr, pubKeyStr, nil
}

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
			Issuer:    issuer,
			Subject:   "user",
			Audience:  []string{audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a token object, specifying the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(privKey.ToECDSA())
}

func VerifyJWTToken(tokenString, pubKeyStr string) (jwt.MapClaims, error) {
	// Decode the base64 public key string
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %w", err)
	}

	// Parse the public key
	pubKey, err := secp256k1.ParsePubKey(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// Verify the token using the parsed public key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return pubKey.ToECDSA(), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	// Return the claims if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
