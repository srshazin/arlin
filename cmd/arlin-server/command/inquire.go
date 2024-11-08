package command

import (
	"errors"

	"github.com/gorilla/websocket"
	appstate "shazin.me/arlin/cmd/arlin-server/app_state"
	"shazin.me/arlin/cmd/shared"
)

// command checks if device is paired then replies with a message that contains device information

type inqData struct {
	DeviceID    string `json:"deviceID"`
	HostName    string `json:"hostName"`
	HostAddress string `json:"hostAddress"`
	Port        int    `json:"port"`
}

func SendINQMsg(deviceId string, conn *websocket.Conn) error {
	// check if the device is paired or not
	devicePaired, _ := appstate.IsDevicePaired(deviceId)

	if !devicePaired {
		conn.WriteMessage(websocket.TextMessage, []byte("Device is not paired!"))
		conn.Close()
		return errors.New("Device not paired")
	}

	deviceIP, error := shared.GetDeviceLocalIP()

	if error != nil {
		return error
	}

	hostName, error := shared.GetDeviceHostName()

	if error != nil {
		return error
	}

	appState, error := appstate.GetAppState()

	if error != nil {
		return error
	}

	inqMsg := inqData{
		DeviceID:    appState.DeviceID,
		HostName:    hostName,
		HostAddress: deviceIP,
		Port:        shared.GetServicePort(),
	}
	conn.WriteJSON(inqMsg)

	return nil
}
