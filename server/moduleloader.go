package server

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/etimo/go-magic-mirror/server/modules"
	"github.com/etimo/go-magic-mirror/server/modules/clock"
	"github.com/etimo/go-magic-mirror/server/modules/systeminfo"
)

//ModuleContext : struct that contains array of all server side modules,
//and connects them with the correct channel for sending messages.
type ModuleContext struct {
	Modules      []modules.Module
	Creators     map[string]moduleCreator
	WriteChannel chan []byte
	ReadChannel  chan []byte
}
type CreateMessage struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type moduleCreator modules.Module

//NewModuleContext creates a moduleContext with default set of modules,
//and module creators.
//The module context contains all modules initiated serverside and connects them
//to the right channels for sending and receiving messsages.
func NewModuleContext(writeChannel chan []byte, readChannel chan []byte) ModuleContext {
	var mods = make([]modules.Module, 0)
	//	mods = append(mods, systeminfo.NewSysInfoModule(writeChannel, "systeminfo", 200*time.Millisecond))
	mods = append(mods, clock.NewClockModule(writeChannel, "clock", 1000*time.Millisecond))
	//	mods = append(mods, systeminfo.NewSysInfoModule(writeChannel, "systeminfo2", 500*time.Millisecond))

	moduleCreator := map[string]moduleCreator{
		"systeminfo": systeminfo.SysinfoModule{},
	}
	return ModuleContext{
		Modules:      mods,
		Creators:     moduleCreator,
		WriteChannel: writeChannel,
		ReadChannel:  readChannel,
	}

}

//SetupTimedUpdates starts the timedUpdate flow for all modules in the module list
//Should only be called once on startup.
func (m ModuleContext) SetupTimedUpdates() {

	for _, module := range m.Modules {
		fmt.Printf("module: %v\n", module)
		go module.TimedUpdate()
	}
}

//RecieveCreateMessage handles incoming messages from the frontend and initiate
//modules on the server. This can be used instead of creating them on the server
//during construction. Message sent from frontend must match the createMessage
//struct and each module places own demands on the inner message.
func (m ModuleContext) RecieveCreateMessage() {
	for {
		incoming := <-m.ReadChannel
		var request CreateMessage

		err := json.Unmarshal(incoming, &request)
		if err != nil {
			continue
		}
		log.Printf("Received createione %v\n", request)
		m.handleMessage(request, incoming)
	}
}
func (m ModuleContext) handleMessage(request CreateMessage, incoming []byte) {

	creator := m.Creators[request.Name]
	if creator == nil {
		return
	}
	for _, mod := range m.Modules {
		//Prevent duplicate module initiations
		if mod.GetId() == request.ID {
			log.Printf("There is already a module with id: %s", request.ID)
			return
		}
	}

	mod, err := creator.CreateFromMessage(incoming, m.WriteChannel)
	if err == nil && mod.GetId() == request.ID {
		m.Modules = append(m.Modules, mod)
		go mod.TimedUpdate()
		log.Printf("Added %v %v!", mod, err)
	}
}

//InitialMessages sends updates for all current modules
//Will be called when a new WS is established to send initial data.
func (m ModuleContext) InitialMessages() {
	for _, module := range m.Modules {
		module.Update()
	}

}
