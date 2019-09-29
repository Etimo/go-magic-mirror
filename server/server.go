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
	router.PathPrefix("/static/").Handler(http.StripPrefix("./dist/",
		http.FileServer(http.Dir("./public"))))
	go socket.WriteWaiting()

	handler := HandleError(
		handlers.LoggingHandler(os.Stdout, router))

	log.Fatal(http.ListenAndServe(bindAddress, handler))

}
