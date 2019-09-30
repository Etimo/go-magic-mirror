package systeminfo

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	channel := make(chan []byte, 20)
	module := NewSysInfoModule(channel, "testmodule", 500*time.Millisecond)
	module.Update()

	returnBytes := <-channel
	testMessage := SysMessage{}
	err := json.Unmarshal(returnBytes, &testMessage)
	fmt.Printf("Received: %+v", testMessage)
	if err != nil {
		log.Fatal("Could not unmarshal message")
		t.Fail()
	}
}
