package main

import (
	"fmt"

	arlinserver "shazin.me/arlin/cmd/arlin-server"
	"shazin.me/arlin/cmd/arlin-server/models"
)

func main() {
	// initApp()
	arlinserver.InitAppStats()
	fmt.Println(arlinserver.GetAppState())
	arlinserver.AddPairedDevice(models.ArlinPairedDeviceInfo{
		DeviceID:   "XEADSIU34SX",
		DeviceName: "iPhone 16 Pro Max",
	})
	fmt.Println(arlinserver.GetAppState())
}
