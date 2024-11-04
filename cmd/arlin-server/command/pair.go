package command

import (
	"encoding/json"
	"fmt"
)

type PairingDevice struct {
	DeviceModel string
}

func PairDevice(conn_data string) error {
	var pairingDevice PairingDevice
	error := json.Unmarshal([]byte(conn_data), &pairingDevice)

	if error != nil {
		return error
	}
	fmt.Println("Connection request from: ", pairingDevice.DeviceModel)

	return nil

}
