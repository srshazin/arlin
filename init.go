package main

import (
	"log"
	"sync"

	arlinmdns "shazin.me/arlin/cmd/arlin-mdns"
	arlinserver "shazin.me/arlin/cmd/arlin-server"
)

// create a channel for holding the port number
var portChannel = make(chan int)

// create a worker group
var wg sync.WaitGroup

func initApp() {
	// first init app state
	error := arlinserver.InitAppStats()

	if error != nil {
		log.Fatal(error)
	}

	// add the websocket to the worker
	wg.Add(1)
	go arlinserver.FireUpWsServer(portChannel)

	// add the mdns service advertising server worker
	wg.Add(1)
	go arlinmdns.Broadcast(portChannel)

	// wait for the goroutines to finish
	wg.Wait()
}
