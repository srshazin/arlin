package arlinmdns

import (
	"fmt"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

func Broadcast() {
	// Define the details of the mDNS service
	serviceName := "Arlin remote linux service"
	serviceType := "_arlin._tcp"
	domain := "local."
	port := 8080

	// Register the service with zeroconf
	server, err := zeroconf.Register(
		serviceName,                  // Service instance name
		serviceType,                  // Service type
		domain,                       // Domain, usually "local."
		port,                         // Port where the service is accessible
		[]string{"txtv=hello world"}, // Optional metadata (TXT records)
		nil,                          // Use default network interface
	)
	if err != nil {
		log.Fatalf("Failed to register mDNS service: %v", err)
	}
	defer server.Shutdown()

	fmt.Printf("Service '%s' is now advertised on the network as '%s' at port %d.\n", serviceName, serviceType, port)

	// Keep the program running to continue broadcasting
	select {
	case <-time.After(10 * time.Minute):
		fmt.Println("Stopping advertisement after 10 minutes.")
	}
}
