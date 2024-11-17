package command

import (
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

func HandleKeyPress(key string, conn *websocket.Conn) error {
	error := robotgo.KeyTap(key)
	return error
}
