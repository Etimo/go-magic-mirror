package server

import (
	"log"
	"net/http"
	"os"

	"github.com/etimo/go-magic-mirror/server/controllers"
	"github.com/etimo/go-magic-mirror/server/recovery"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartServer(bindAddress string) {
	router := mux.NewRouter()
	router.HandleFunc("/api/pomodoro", controllers.PomodoroReturn)
	router.HandleFunc("/panictest", func(w http.ResponseWriter, r *http.Request) {
		panic("This is a triggered panic")
	})
	router.PathPrefix("/").Handler(http.StripPrefix("/",
		http.FileServer(http.Dir("./dist"))))
	//Handlers are methods called on all routes they are registered for,
	//here we register a LoggingHandler for access-tracking and a recovery (crash-handler)
	handler := recovery.HandleRecovery(
		handlers.LoggingHandler(os.Stdout, router))
	log.Fatal(http.ListenAndServe(bindAddress, handler))

}
