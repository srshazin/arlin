package arlinmdns

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/grandcat/zeroconf"
)

func Broadcast(ch <-chan int) {
	port := <-ch
	// Define the details of the mDNS service
	hostAddress, _ := getLocalIP()
	serviceName := base64.StdEncoding.EncodeToString([]byte(hostAddress))
	serviceType := "_arlin._tcp"
	domain := "local."
	fmt.Println("Service: ", serviceName)
	hostName, _ := os.Hostname()

	// Register the service with zeroconf
	server, err := zeroconf.Register(
		serviceName, // Service instance name
		serviceType, // Service type
		domain,      // Domain, usually "local."
		port,        // Port where the service is accessible
		[]string{fmt.Sprintf("host=%s", hostName)}, //  metadata
		nil, // Use default network interface
	)
	if err != nil {
		log.Fatalf("Failed to register mDNS service: %v", err)
	}
	defer server.Shutdown()

	fmt.Printf("Service '%s' is now advertised on the network as '%s' at port %d.\n", serviceName, serviceType, port)

	// Keep the program running to continue broadcasting
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Stopping advertisement after 10 minutes.")
	}
}
