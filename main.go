package main

import (
	"fmt"

	arlinserver "shazin.me/arlin/cmd/arlin-server"
)

func main() {
	// initApp()
	arlinserver.InitAppStats()
	fmt.Println(arlinserver.GetAppState())
}
