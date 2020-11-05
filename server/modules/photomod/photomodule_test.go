package photomod

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func testPhotoModule(t *testing.T) {

	channel := make(chan []byte, 20)
	module := NewPhotoModule(channel, "testmodule", "URL", 500*time.Millisecond)
	module.Update()
	returnBytes := <-channel
	testMessage := PhotoMessage{}
	err := json.Unmarshal(returnBytes, &testMessage)
	fmt.Printf("Received: %+v", testMessage)
	if err != nil {
		log.Fatal("Could not unmarshal message")
		t.Fail()
	}
	if testMessage.Url != "URL" {
		log.Fatal("Received bad message")
		t.Fail()

	}
}
