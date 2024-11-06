package command

import (
	"encoding/json"
	"fmt"

	"shazin.me/arlin/cmd/arlin-server/utils"
)

type PairingDevice struct {
	DeviceModel string
	Brand       string
	DeviceID    string
}

func PairDevice(conn_data string) error {
	var pairingDevice PairingDevice
	error := json.Unmarshal([]byte(conn_data), &pairingDevice)

	if error != nil {
		return error
	}
	accepted, error := utils.PromptLinux(fmt.Sprintf("Device %s %s with id %s is asking to pair. Do you accept?", pairingDevice.Brand, pairingDevice.DeviceModel, pairingDevice.DeviceID))

	if accepted {
		fmt.Println("Connected to device")
	} else {
		fmt.Println("Connection rejected!")
	}
	return nil

}
