package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
	currentWs       *websocket.Conn
	WriteChannel    chan []byte //This channel should be collected by other components.
	ReadChannel     chan []byte
	OpChannel       chan []byte
	CallbackChannel chan bool
}

//NewServerSocket : Will create a new ServerSocket that can listen for connections.
//callback : This function will be called every time a socket connection is established
func NewServerSocket(callbackChannel chan bool) ServerSocket {
	return ServerSocket{
		WriteChannel:    make(chan []byte, 20),
		ReadChannel:     make(chan []byte, 20),
		CallbackChannel: callbackChannel,
	}
}

//BindWebSocket : This method can be called from a Go http router, it will upgrade the connection
//to a websocket and store it.
//Currently allows one connection, reconnecting closes existing connection.
func (s *ServerSocket) BindWebSocket(w http.ResponseWriter, r *http.Request) {

	log.Println("Establishing connection")
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not upgrade connection %v\n", r))
		return
	}

	if s.currentWs != nil {
		log.Println("Killing old WS connection")
		s.currentWs.Close() //Clean slate
		s.currentWs = nil
	}

	log.Println("Linking new WS connection")
	s.currentWs = connection
	defer s.ReadIncoming(connection)

	data, _ := json.Marshal(models.LayoutMessage{Rows: 4, Cols: 6})
	log.Println("Write to WS connection")
	s.WriteChannel <- data
	s.callbackTrigger()
	log.Println("Connected new websocket")

	// defer func() {
	// 	time.Sleep(10 * time.Second)
	// 	data, _ := json.Marshal(models.LayoutMessage{Rows: 2, Cols: 2})
	// 	log.Println("Write to WS connection")
	// 	s.WriteChannel <- data
	// 	s.callbackTrigger()
	// }()
}
func (s *ServerSocket) callbackTrigger() {
	if s.CallbackChannel != nil {
		log.Println("Running on connection callback")
		select {
		case s.CallbackChannel <- true:
			log.Println("Triggered callback channel")
		default:
			log.Println("Callback channel is full")
		}
	}
}

//ReadIncoming : Reads messages incoming as byte arrays from the websocket
//and writes them to the ReadChannel.
func (s *ServerSocket) ReadIncoming(ws *websocket.Conn) {
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading from socket")
			break
		}
		s.ReadChannel <- message
		fmt.Println("Received message", messageType, " : ", len(message))
	}
	log.Println("Ended readloop")
}
func (s *ServerSocket) WriteWaiting() {
	for {
		//Comment!
		if s.currentWs == nil {
			fmt.Printf("Socket still nill!")
			time.Sleep(500 * time.Millisecond)
			continue
		}
		writeByte := <-s.WriteChannel

		if isJSON(writeByte) {
			//fmt.Println("Message written: ", string(writeByte))
			err := s.currentWs.WriteMessage(websocket.TextMessage, writeByte)
			if err != nil {
				fmt.Printf("Err writing to socket: %v\n", err)
			}
		} else {
			fmt.Println("Not valid JSON: ", string(writeByte))
		}

	}
}

func isJSON(bytes []byte) bool {
	var js json.RawMessage
	err := json.Unmarshal(bytes, &js)
	if err != nil {
		fmt.Println("ERROR HERE!", err == nil, " ", err.Error())
	}
	return err == nil
}
