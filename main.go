package main

import (
	"fmt"
	"log"

	arlinserver "shazin.me/arlin/cmd/arlin-server"
)

func main() {
	// initApp()
	arlinserver.InitAppStats()
	fmt.Println(arlinserver.GetAppState())
	// error := arlinserver.AddPairedDevice(models.ArlinPairedDeviceInfo{
	// 	DeviceID:   "XEADSIU34SX",
	// 	DeviceName: "iPhone 16 Pro Max",
	// })
	//
	error := arlinserver.UnpairDevice("XEU34SX")

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(arlinserver.GetAppState())
}
