package command

import (
	"github.com/gorilla/websocket"
	appstate "shazin.me/arlin/cmd/arlin-server/app_state"
)

func connectDevice(deviceID string, conn *websocket.Conn) error {
	// check whether device is already paired
	devicePaired, _ := appstate.IsDevicePaired(deviceID)
	if devicePaired {
		error := conn.WriteMessage(websocket.TextMessage, []byte("OK"))
		return error
	} else {
		error := conn.WriteMessage(websocket.TextMessage, []byte("UNPAIRED"))
		conn.Close()
		return error

	}
}
