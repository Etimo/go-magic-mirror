package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ServerSocket struct {
	currentWs         *websocket.Conn
	WriteChannel      chan []byte //This channel should be collected by other components.
	ConnectedCallback func()
}

func NewServerSocket(callback func()) ServerSocket {
	return ServerSocket{
		WriteChannel:      make(chan []byte, 20),
		ConnectedCallback: callback,
	}
}

//Currently allows one connection, killing the old one.
//Needs pointer reference to associated socket field.
func (s *ServerSocket) BindWebSocket(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not upgrade connection %v\n", r))
		return
	}

	if s.currentWs != nil {
		log.Println("Killing old WS connection")
		s.currentWs.Close() //Clean slate
	}

	log.Println("Linking new WS connection")
	s.currentWs = connection

	json, _ := json.Marshal(models.WelcomeMessage{Message: "Connected socket.."})
	log.Println("Write to WS connection")
	s.WriteChannel <- json
	if s.ConnectedCallback != nil {
		s.ConnectedCallback()
	}
}

func (s *ServerSocket) WriteWaiting() {
	for {
		//Comment!
		writeByte := <-s.WriteChannel
		if s.currentWs == nil {
			fmt.Printf("Socket still nill!")
			continue
		}
		err := s.currentWs.WriteMessage(websocket.TextMessage, writeByte)
		if err != nil {
			fmt.Printf("Err writing to socket: %v\n", err)
		}
	}

}
