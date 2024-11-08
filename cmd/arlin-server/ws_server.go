package arlinserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"shazin.me/arlin/cmd/arlin-server/command"
	"shazin.me/arlin/cmd/shared"
)

// Upgrader for HTTP to WS Protocol
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Handler function for websocket
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error connecting to ws server: ", err)
		return
	}

	defer conn.Close()

	// An Inifinite loop to listen for connection
	for {
		_, msg, error := conn.ReadMessage()

		if error != nil {
			log.Println("Client disconnected")
			conn.Close()
			break
		}
		cmd := string(msg)
		parsedCmd, error := command.ParseCommand(cmd)
		// fmt.Println(parsedCmd)
		command.ExecuteCommand(parsedCmd, conn)
	}

}

func FireUpWsServer(ch chan<- int) {
	http.HandleFunc("/ws", wsHandler)
	port := shared.GetServicePort()
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("ws server started on port %v\n", port)
	ch <- port
	log.Fatal(http.ListenAndServe(addr, nil))
}
