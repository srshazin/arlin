package utils

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// Function to select a random available port
func GetRandomPort() int {
	rand.Seed(time.Now().UnixNano())
	for {
		port := rand.Intn(10000) + 10000 // Random port between 10000 and 20000
		if isPortAvailable(port) {
			return port
		}
	}
}

// Check if the port is available
func isPortAvailable(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	ln.Close()
	return true
}
