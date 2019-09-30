package server

import (
	"log"
	"net/http"
	"os"

	"github.com/etimo/go-magic-mirror/server/controllers"
	"github.com/etimo/go-magic-mirror/server/modules"
	"github.com/etimo/go-magic-mirror/server/socket"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var mods modules.ModuleContext
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
	router.HandleFunc("/api/pomodoro", contrl.PomodoroReturn)
	router.HandleFunc("/ws", sock.BindWebSocket)
	router.HandleFunc("/forward", contrl.WriteToChannel)
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
	SetupModules()
	sock.ConnectedCallback = mods.InitialMessages
	log.Fatal(http.ListenAndServe(bindAddress, handler))

}
func SetupModules() {
	mods = modules.NewModuleContext(sock.WriteChannel)
	go mods.SetupTimedUpdates()
}
