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
		DeviceID:   "XEU34SX",
		DeviceName: "Google Pixel 4A",
	})
	fmt.Println(arlinserver.GetAppState())
}
