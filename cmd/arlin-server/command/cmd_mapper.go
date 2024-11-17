package command

import (
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
)

// ExecuteCommand executes a parsed command based on its action and parameters
func ExecuteCommand(cmd *Command, conn *websocket.Conn) {
	switch cmd.Action {
	case "MOVE":
		// fmt.Printf("Moving to coordinates x=%s, y=%s\n", cmd.Params["dx"], cmd.Params["dy"])
		dx := cmd.Params["dx"]
		dy := cmd.Params["dy"]
		dxInt, _ := strconv.Atoi(dx)
		dyInt, _ := strconv.Atoi(dy)
		moveCursor(dxInt, dyInt)
	case "PRESS":
		HandleKeyPress(cmd.Params["key"], conn)
	case "MOUSE":
		handleMouseClick(cmd.Params["button"])
	case "CONNECT":
		connectDevice(cmd.Params["deviceID"], conn)
	case "PAIR":
		PairDevice(cmd.Params["data"], conn)
	case "INQ":
		SendINQMsg(cmd.Params["deviceID"], conn)
	default:
		fmt.Println("Unknown command:", cmd.Action)
	}
}
