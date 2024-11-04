package main

import (
	"log"

	arlinserver "shazin.me/arlin/cmd/arlin-server"
)

func main() {
	// initApp()
	error := arlinserver.InitAppStats()
	if error != nil {
		log.Fatal(error)
	}
}
