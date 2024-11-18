package command

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	appstate "shazin.me/arlin/cmd/arlin-server/app_state"
	"shazin.me/arlin/cmd/arlin-server/models"
	"shazin.me/arlin/cmd/arlin-server/utils"
)

type PairingDevice struct {
	DeviceModel string
	Brand       string
	DeviceID    string
}

func PairDevice(conn_data string, conn *websocket.Conn) error {
	fmt.Println("Echo from pairing")
	var pairingDevice PairingDevice
	error := json.Unmarshal([]byte(conn_data), &pairingDevice)

	if error != nil {
		fmt.Println(error)
		return error
	}
	accepted, error := utils.PromptLinux(fmt.Sprintf("Device %s %s with id %s is asking to pair. Do you accept?", pairingDevice.Brand, pairingDevice.DeviceModel, pairingDevice.DeviceID))
	if error != nil {
		fmt.Println(error)
		return error
	}
	if accepted {
		fmt.Println("Connected to device")
		conn.WriteMessage(websocket.TextMessage, []byte("PAIRING_ACCEPTED"))
		appstate.AddPairedDevice(models.ArlinPairedDeviceInfo{
			DeviceID:   pairingDevice.DeviceID,
			DeviceName: pairingDevice.DeviceModel,
		})
	} else {

		fmt.Println("Connection rejected!")
		fmt.Printf("Prompt stat:%v and %v\n", accepted, error)
		conn.WriteMessage(websocket.TextMessage, []byte("PAIRING_REJECTED"))
		conn.Close()
	}
	return nil

}
