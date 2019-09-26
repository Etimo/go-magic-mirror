package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/etimo/go-magic-mirror/server/models"
)

func PomodoroReturn(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Pomodoro{Simple: "Pomodoro was here!"})
}
