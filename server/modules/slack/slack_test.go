package slackmodule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/slack-go/slack"
)

var fakeSlackSourceMessage slack.Message = slack.Message{
	Msg: slack.Msg{
		Username:  "ERIK",
		Timestamp: "1234",
		Text:      "JAG ÄR SÅ BALL",
	},
}

type testEventSource struct {
}

type FakeSlackProvider struct{}

func (FakeSlackProvider) GetLatestMessages(max int) []slack.Message {
	return []slack.Message{
		fakeSlackSourceMessage,
	}
}

func TestFakeSlackFetch(t *testing.T) {
	fmt.Println("ACTUAL!")
	//Capture stdout
	var buf bytes.Buffer
	log.SetOutput(&buf)

	writeChannel := make(chan []byte, 100)

	module := NewSlackModule(
		writeChannel,
		"testSlack",
		time.Duration(500),
		time.Duration(500),
		FakeSlackProvider{})

	fmt.Printf("Mod: %v\n", module)
	module.Update()

	var message SlackUpdateMessage
	err := json.Unmarshal(<-writeChannel, &message)
	if err != nil {
		log.Println("Failed tes, could not unmarshal output: ", err.Error())
		t.Fail()
	}
	if !reflect.DeepEqual(message.SlackMessages[0], fakeSlackSourceMessage) {
		log.Println("Message returned from slack-module not expected one")

	}

	fmt.Println("Events: ", message)
	fmt.Println("Stdout: ", buf.String())
}

//This test demands active slack token, and connectivity
//so is commented out. Leaving this in to allow easy testing of credential setup.
func TestActualSlackFetch(t *testing.T) {
	//Disabled test
	if true {
		return
	}
	fmt.Println("ACTUAL!")
	//Capture stdout
	var buf bytes.Buffer
	log.SetOutput(&buf)

	writeChannel := make(chan []byte, 100)
	module := NewSlackModule(
		writeChannel,
		"testSlack",
		time.Duration(500),
		time.Duration(500),
		GetSlackProvider(
			os.Getenv("SlackTestToken"),
			"etimo_internal"))
	fmt.Printf("Mod: %v\n", module)
	module.Update()
	module.Update()

	var message SlackUpdateMessage
	err := json.Unmarshal(<-writeChannel, &message)
	if err != nil {
		log.Println("Failed tes, could not unmarshal output: ", err.Error())
		t.Fail()
	}

	fmt.Println("Events: ", message)
	fmt.Println("Stdout: ", buf.String())
}
