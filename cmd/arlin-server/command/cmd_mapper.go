package command

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// ExecuteCommand executes a parsed command based on its action and parameters
func ExecuteCommand(cmd *Command, conn *websocket.Conn) {
	switch cmd.Action {
	case "MOVE":
		fmt.Printf("Moving to coordinates x=%s, y=%s\n", cmd.Params["x"], cmd.Params["y"])
	case "PRESS":
		fmt.Printf("Pressing key: %s\n", cmd.Params["key"])
	case "MOUSE":
		fmt.Printf("Clicking %s button\n", cmd.Params["button"])
	case "CONNECT":
		fmt.Printf("Connecting to deviceID: %s\n", cmd.Params["deviceID"])
	case "PAIR":
		PairDevice(cmd.Params["data"], conn)
	case "INQ":
		SendINQMsg(cmd.Params["deviceID"], conn)
	default:
		fmt.Println("Unknown command:", cmd.Action)
	}
}
