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

type Server struct {
	mods   *ModuleContext
	sock   *socket.ServerSocket
	contrl controllers.Controllers
}

func (s *Server) baseSetup() {
	sock := socket.NewServerSocket(make(chan bool, 20))
	s.sock = &sock
	s.contrl = controllers.Controllers{
		SocketChannel: s.sock.WriteChannel,
	}
	go s.sock.WriteWaiting()

}
func StartServer(bindAddress string) {
	s := Server{}
	s.baseSetup()

	router := mux.NewRouter()
	router.Handle("/public", http.FileServer(http.Dir("./public")))

	router.HandleFunc("/ws", s.sock.BindWebSocket)
	router.HandleFunc("/forward", s.contrl.WriteToChannel).Methods(http.MethodPost)

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
	s.setupModules()
	log.Fatal(http.ListenAndServe(bindAddress, handler))

}

func (s *Server) setupModules() {
	mods := NewModuleContext(s.sock.WriteChannel,
		s.sock.ReadChannel,
		s.sock.CallbackChannel)
	s.mods = &mods
	s.sock.CallbackChannel = s.mods.CallbackChannel
	go mods.SetupTimedUpdates()
	go RecieveCreateMessage(s.mods)
	go ReadCallback(s.mods)
}
