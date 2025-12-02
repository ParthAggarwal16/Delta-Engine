package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateID creates a random 16-byte hex string to use as a unique ID
func GenerateID() string {
	b := make([]byte, 16) // 16 bytes = 128 bits
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error generating ID:", err)
		return ""
	}
	return hex.EncodeToString(b)
}
