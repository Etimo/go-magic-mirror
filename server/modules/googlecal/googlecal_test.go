package googlecal

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type testEventSource struct {
}

var testMessage UpdateMessage = UpdateMessage{
	CalendarName: "Fake calendar",
	Events:       []EventMessage{EventMessage{}},
}

func (gc testEventSource) GetEvents(
	startDateTime string, stopDateTime string, numberOfEvents int, initialLoad bool) []UpdateMessage {
	returnMessages := make([]UpdateMessage, 1)
	return append(returnMessages, testMessage)
}

//TestSimpleSource tests using a faked event source backend.
func TestSimpleSource(t *testing.T) {
	writeChannel := make(chan []byte, 100)
	module := NewGoogleCalendarModuleAlternativeSource(testEventSource{}, "TestModule", writeChannel)
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

//This test demands active google-calendar credentials, and connectivity
//so is commented out. Leaving this in to allow easy testing of credential setup.
/*
func TestActualUpdate(t *testing.T) {
	writeChannel := make(chan []byte, 100)
	module, error := NewGoogleCalendarModule(
		"testCal", []string{"Etvrimo Event-bokning"},
		writeChannel,
	)
	if error != nil {
		log.Println("Failed test, could not connect to google.")
		t.Fail()
	}
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
*/
