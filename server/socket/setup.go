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

//ServerSocket : Represents a socket connection to the backend, will only allow one concurrent connection.
type ServerSocket struct {
	currentWs         *websocket.Conn
	WriteChannel      chan []byte //This channel should be collected by other components.
	ReadChannel       chan []byte
	ConnectedCallback func()
}

//NewServerSocket : Will create a new ServerSocket that can listen for connections.
//callback : This function will be called every time a socket connection is established
func NewServerSocket(callback func()) ServerSocket {
	return ServerSocket{
		WriteChannel:      make(chan []byte, 20),
		ReadChannel:       make(chan []byte, 20),
		ConnectedCallback: callback,
	}
}

//BindWebSocket : This method can be called from a Go http router, it will upgrade the connection
//to a websocket and store it.
//Currently allows one connection, reconnecting closes existing connection.
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

//ReadIncoming : Reads messages incoming as byte arrays from the websocket
//and writes them to the ReadChannel.
func (s *ServerSocket) ReadIncoming() {
	for {
		messageType, message, err := s.currentWs.ReadMessage()
		if err != nil {
			log.Fatal("Error reading from socket")
			continue
		}
		fmt.Println("Received create message", messageType, " : ", len(message))
		s.ReadChannel <- message
	}
}
func (s *ServerSocket) WriteWaiting() {
	for {
		//Comment!
		writeByte := <-s.WriteChannel
		if s.currentWs == nil {
			//fmt.Printf("Socket still nill!")
			continue
		}
		err := s.currentWs.WriteMessage(websocket.TextMessage, writeByte)
		if err != nil {
			fmt.Printf("Err writing to socket: %v\n", err)
			_ = s.currentWs.Close()
			s.currentWs = nil
		}
	}

}
