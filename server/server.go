package server

import (
	"log"
	"net/http"
	"os"

	"github.com/etimo/go-magic-mirror/server/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.Handle("/public", http.FileServer(http.Dir("./public")))
	router.HandleFunc("/api/pomodoro", controllers.PomodoroReturn)
	router.PathPrefix("/static/").Handler(http.StripPrefix("./dist/",
		http.FileServer(http.Dir("./public"))))

	handler := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe("localhost:8080", handler))

}
