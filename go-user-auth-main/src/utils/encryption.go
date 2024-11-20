package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// add function to encryp using sha256 the password based on the salt provided
func Encrypt(password string) string {
	h := sha256.New()
	h.Write([]byte(password))

	// return the hex encoded string
	return hex.EncodeToString(h.Sum(nil))
}

// add function to compare the password and the hash
func Compare(password, hash string) bool {
	h := sha256.New()
	h.Write([]byte(password))

	// return the comparison
	return hex.EncodeToString(h.Sum(nil)) == hash
}
