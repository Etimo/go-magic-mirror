package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/gorilla/websocket"
)

var currentWs *websocket.Conn
var WriteChannel = make(chan []byte) //This channel should be collected by other components.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Currently allows one connection, killing the old one.
func BindWebSocket(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not upgrade connection %v\n", r))
		return
	}

	if currentWs != nil {
		log.Println("Killing old WS connection")
		currentWs.Close() //Clean slate
	}

	log.Println("Linking new WS connection")
	currentWs = connection

	json, _ := json.Marshal(models.WelcomeMessage{Message: "Connected socket.."})
	WriteChannel <- json
}

func WriteWaiting() {
	fmt.Printf("Starting socket writeloop")
	for {
		//Comment!
		writeByte := <-WriteChannel
		if currentWs == nil {
			continue
		}
		err := currentWs.WriteMessage(websocket.TextMessage, writeByte)
		if err != nil {
			fmt.Printf("Err writing to socket: %v\n", err)
		}
		fmt.Printf("Socket wroteloop!")
	}

}
