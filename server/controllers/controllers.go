package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/etimo/go-magic-mirror/server/socket"
)

var channelWriter = socket.WriteChannel

func PomodoroReturn(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Pomodoro{Simple: "Pomodoro was here!"})
}

func WriteToChannel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Fatal("API supports only POST, was:", r.Method)
		http.Error(w, "Wrong method", http.StatusBadRequest)
	}

	body, err := r.GetBody()
	if err != nil {
		log.Fatal("Error when calling writeToChannel", err)
		http.Error(w, "No body in call.", http.StatusBadRequest)
	}
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal("Error when calling writeToChannel", err)
		http.Error(w, "Error reading body", http.StatusBadRequest)
	}
	//This pushes incoming bytes to our websocket for easy testing
	channelWriter <- bytes

}
