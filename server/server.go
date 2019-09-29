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

func StartServer(bindAddress string) {
	router := mux.NewRouter()
	router.Handle("/public", http.FileServer(http.Dir("./public")))
	router.HandleFunc("/api/pomodoro", controllers.PomodoroReturn)
	router.HandleFunc("/ws", socket.BindWebSocket)
	router.HandleFunc("/forward", controllers.WriteToChannel)
	router.PathPrefix("/").Handler(http.StripPrefix("/",
		http.FileServer(http.Dir("./dist"))))
	//Handlers are methods called on all routes they are registered for,
	//here we register a LoggingHandler for access-tracking and a recovery (crash-handler)
	go socket.WriteWaiting()

	handler := HandleError(
		handlers.LoggingHandler(os.Stdout, router))
	router.HandleFunc("/panictest", func(w http.ResponseWriter, r *http.Request) {
		panic("This is a triggered panic")
	})
	log.Fatal(http.ListenAndServe(bindAddress, handler))

}
