package googlecal

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

//This test demands active google-calendar credentials, and connectivity
//so is commented out. Leaving this in to allow easy testing of credential setup.
func TestActualUpdate(t *testing.T) {
	writeChannel := make(chan []byte, 100)
	module := NewGoogleCalendarModule(
		//		os.Getenv("MAGIC_MIRROR_SERVICE_LOCATION"),
		"testCal", []string{"Etvrimo Event-bokning"},
		writeChannel,
	)
	fmt.Printf("Mod: %v\n", module)
	module.Update()
	var message Message
	err := json.Unmarshal(<-writeChannel, &message)
	if err != nil {
		log.Println("Failed tes, could not unmarshal output: ", err.Error())
		t.Fail()
	}
	log.Println("Events: ", message)
}
