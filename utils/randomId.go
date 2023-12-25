package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomID() (string, error) {
	length := 10
	// Calculate the number of bytes needed
	numBytes := length / 2

	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Convert bytes to hexadecimal string
	randomID := hex.EncodeToString(randomBytes)

	// Trim the string to the desired length
	randomID = randomID[:length]

	return randomID, nil
}
