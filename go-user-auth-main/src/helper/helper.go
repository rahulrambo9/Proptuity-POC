package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	base64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"math/big"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
)

const (
	NoPadding Padding = iota
)

const (
	GCM CipherMode = iota
)

type CipherMode int
type Padding int

const ENCRYPT = "encrypt"
const DECRYPT = "decrypt"

type AES struct {
	CipherMode CipherMode
	Padding    Padding
}

func EcdsaHelper(method, text string, key []byte) (string, error) {
	if method == ENCRYPT {
		return Encrypt(text, key)
	} else if method == DECRYPT {
		return Decrypt(text, key)
	} else {
		return "", fmt.Errorf("EcdsaHelper invalid method %s", method)
	}
}

func Encrypt(plainText string, key []byte) (string, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//block cipher wrapped in Galois Counter Mode
	gcm, err := cipher.NewGCMWithNonceSize(aes, 16)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	plainTextBytes := []byte(plainText)
	cipherText := gcm.Seal(nil, nonce, plainTextBytes, nil)
	return packCipherData(cipherText, nonce, gcm.Overhead()), nil
}

func Decrypt(cipherText string, key []byte) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	aes, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//block cipher wrapped in Galois Counter Mode
	aesgcm, err := cipher.NewGCMWithNonceSize(aes, 16)
	if err != nil {
		return "", err
	}
	encryptedBytes, nonce := unpackCipherData(encrypted, aesgcm.NonceSize())
	decryptedBytes, err := aesgcm.Open(nil, nonce, encryptedBytes, nil)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes[:]), nil
}

// func GenerateKeyPair() (string, string, error) {
// 	privKey, err := secp256k1.GeneratePrivateKey()
// 	if err != nil {
// 		return "", "", err
// 	}
// 	pubKeyBytes := privKey.PubKey().SerializeUncompressed()
// 	dBytes := privKey.ToECDSA().D.Bytes()
// 	return base64.StdEncoding.EncodeToString(dBytes),
// 		base64.StdEncoding.EncodeToString(pubKeyBytes[1:]), err
// }

func GenerateKeyPairK1() (string, string, error) {
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

func GenerateKeyPair() (string, string, error) {
	// Use the P-256 curve (secp256r1)
	curve := elliptic.P256()

	// Generate private key
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return "", "", err
	}

	// Encode private key as base64
	privKeyStr := base64.StdEncoding.EncodeToString(privKey.D.Bytes())

	// Encode public key, including the 0x04 prefix for uncompressed format
	pubKeyBytes := elliptic.Marshal(curve, privKey.PublicKey.X, privKey.PublicKey.Y)
	pubKeyStr := base64.StdEncoding.EncodeToString(pubKeyBytes)

	return privKeyStr, pubKeyStr, nil
}

func CreateSharedKey(privKeyStr string, pubKeyStr string) ([]byte, error) {
	// Decode private and public keys from base64
	privKeyBytes, err := base64.StdEncoding.DecodeString(privKeyStr)
	if err != nil {
		return nil, fmt.Errorf("CreateSharedKey error decoding private key: %v", err)
	}
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		return nil, fmt.Errorf("CreateSharedKey error decoding public key: %v", err)
	}

	fmt.Println("Public Key : ", (pubKeyBytes))
	fmt.Println("Private Key : ", (privKeyBytes))

	fmt.Println("Private Key Length: ", len(privKeyBytes))

	// Print key length for debugging
	fmt.Printf("Public Key Length: %d\n", len(pubKeyBytes))

	// Create the private key object
	privKey := new(ecdsa.PrivateKey)
	privKey.D = new(big.Int).SetBytes(privKeyBytes)
	privKey.PublicKey.Curve = elliptic.P256()

	// The private key needs to have the public key set
	privKey.PublicKey.X, privKey.PublicKey.Y = elliptic.P256().ScalarBaseMult(privKey.D.Bytes())

	// Check if the public key is compressed (33 bytes) or uncompressed (65 bytes)
	if len(pubKeyBytes) == 33 {
		// If it's compressed (33 bytes), expand it to uncompressed (65 bytes)
		x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pubKeyBytes)
		if x == nil || y == nil {
			return nil, fmt.Errorf("CreateSharedKey: invalid compressed public key")
		}
		// Reassign the public key bytes as uncompressed
		pubKeyBytes = elliptic.Marshal(elliptic.P256(), x, y)
	} else if len(pubKeyBytes) != 65 {
		return nil, fmt.Errorf("CreateSharedKey: invalid public key format")
	}

	// Unmarshal the public key (expecting uncompressed format now)
	x, y := elliptic.Unmarshal(elliptic.P256(), pubKeyBytes)
	if x == nil || y == nil {
		return nil, fmt.Errorf("CreateSharedKey: invalid public key")
	}
	pubKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	// Generate the shared secret using ECDH (Elliptic Curve Diffie-Hellman)
	sharedSecretX, _ := pubKey.ScalarMult(pubKey.X, pubKey.Y, privKey.D.Bytes())

	// Convert the shared secret to a byte array and optionally apply SHA-256 for key derivation
	sharedKey := sha256.Sum256(sharedSecretX.Bytes())

	return sharedKey[:], nil
}

func CreateSharedKeyK1(privKeyStr string, pubKeyStr string) []byte {
	pubBytes, _ := base64.StdEncoding.DecodeString(pubKeyStr)
	//make into uncompressed format
	tot := make([]byte, len(pubBytes)+1)
	tot[0] = 0x04
	copy(tot[1:], pubBytes)
	pubKey, err := secp256k1.ParsePubKey(tot)
	if err != nil {
		log.Printf("CreateSharedKey error: %v %v", err, len(pubBytes))
	}
	privBytes, _ := base64.StdEncoding.DecodeString(privKeyStr)
	privKey := secp256k1.PrivKeyFromBytes(privBytes)
	return secp256k1.GenerateSharedSecret(privKey, pubKey)
}

// private functions
func packCipherData(cipherText []byte, iv []byte, tagSize int) string {
	ivLen := len(iv)
	data := make([]byte, len(cipherText)+ivLen)
	copy(data[:], iv[0:ivLen])
	copy(data[ivLen:], cipherText)
	return base64.StdEncoding.EncodeToString(data)
}

func unpackCipherData(data []byte, ivSize int) ([]byte, []byte) {
	iv, encryptedBytes := data[:ivSize], data[ivSize:]
	return encryptedBytes, iv
}
