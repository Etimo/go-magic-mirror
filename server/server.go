package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/etimo/go-magic-mirror/server/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var currentWs *websocket.Conn
var WriteChannel = make(chan []byte) //This channel should be collected by other components.

//Currently allows one connection, killing the old one.
func bindWebSocket(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not update connection %v\n", r))
		return
	}

	if connection != nil {
		log.Println("Killing old WS connection")
		connection.Close() //Clean slate
	}

	log.Println("Linking new WS connection")
	currentWs = connection
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func writeWaiting() {
	for {
		//Comment!
		writeByte := <-WriteChannel
		currentWs.WriteMessage(websocket.TextMessage, writeByte)
	}

}
func StartServer(bindAddress string) {
	router := mux.NewRouter()
	router.Handle("/public", http.FileServer(http.Dir("./public")))
	router.HandleFunc("/api/pomodoro", controllers.PomodoroReturn)
	router.PathPrefix("/static/").Handler(http.StripPrefix("./dist/",
		http.FileServer(http.Dir("./public"))))

	handler := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(bindAddress, handler))

}
