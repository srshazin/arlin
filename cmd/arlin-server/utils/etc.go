package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateDeviceID(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random string: %w", err)
		}
		b[i] = charset[randomIndex.Int64()]
	}
	return string(b), nil
}
