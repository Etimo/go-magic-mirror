package server

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/etimo/go-magic-mirror/server/modules"
)

type TestModule struct {
	id string
}

func (tm TestModule) Update() {
}
func (tm TestModule) TimedUpdate() {
}
func (tm TestModule) GetId() string {
	return tm.id
}

type TestCreationMessage struct {
	Id string
}

func (tm TestModule) CreateFromMessage(bytes []byte, channel chan []byte) (modules.Module, error) {
	var mess TestCreationMessage
	err := json.Unmarshal(bytes, &mess)
	if err != nil {
		return nil, err
	}
	fmt.Printf("BLESS THIS MESS: %v\n", mess)
	return TestModule{
		id: mess.Id,
	}, nil
}

func setupContext() *ModuleContext {
	writeChannel := make(chan []byte, 5000)
	readChannel := make(chan []byte, 5000)
	context := NewModuleContext(writeChannel, readChannel)
	context.Modules = make([]modules.Module, 0)
	makers := map[string]moduleCreator{
		"test": TestModule{},
	}
	context.Creators = makers

	return &context
}

func TestCreationViaMessage(t *testing.T) {
	context := setupContext()
	createId := "testId"
	jsonCreate, _ := json.Marshal(TestCreationMessage{
		Id: createId,
	})
	message := CreateMessage{
		Name: "test",
		ID:   createId,
	}

	context.handleMessage(message, jsonCreate)
	fmt.Println("MODULES! ", len(context.Modules))
	if len(context.Modules) == 0 {
		log.Fatal("Did not create module")
		t.Fail()
	}
	if context.Modules[0].GetId() != createId {
		log.Fatal("Did not create with correct id: ", context.Modules[0].GetId(), " : expected ", createId)
		t.Fail()
	}
	context.handleMessage(message, jsonCreate)
	if len(context.Modules) != 1 {
		log.Fatal("Created module with duplicate id")
		t.Fail()
	}
}
