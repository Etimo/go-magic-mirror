package server

import (
	"log"
	"net/http"
	"os"

	"github.com/etimo/go-magic-mirror/server/controllers"
	"github.com/etimo/go-magic-mirror/server/socket"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var mods ModuleContext
var sock socket.ServerSocket
var contrl controllers.Controllers

func baseSetup() {

	sock = socket.NewServerSocket(nil)
	contrl = controllers.Controllers{
		SocketChannel: sock.WriteChannel,
	}
	go sock.WriteWaiting()

}
func StartServer(bindAddress string) {
	baseSetup()

	router := mux.NewRouter()
	router.Handle("/public", http.FileServer(http.Dir("./public")))

	router.HandleFunc("/ws", sock.BindWebSocket)
	router.HandleFunc("/forward", contrl.WriteToChannel).Methods(http.MethodPost)

	router.PathPrefix("/").Handler(
		http.StripPrefix("/",
			http.FileServer(http.Dir("./dist"))))
	//Handlers are methods called on all routes they are registered for,
	//here we register a LoggingHandler for access-tracking and a recovery (crash-handler)

	handler := HandleError(
		handlers.LoggingHandler(os.Stdout, router))
	router.HandleFunc("/panictest", func(w http.ResponseWriter, r *http.Request) {
		panic("This is a triggered panic")
	})
	setupModules()
	sock.ConnectedCallback = mods.InitialMessages
	log.Fatal(http.ListenAndServe(bindAddress, handler))

}
func setupModules() {
	mods = NewModuleContext(sock.WriteChannel, sock.ReadChannel)
	go mods.SetupTimedUpdates()
	go mods.RecieveCreateMessage()
}
