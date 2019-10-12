package server

import (
	"encoding/json"
	"fmt"

	"github.com/etimo/go-magic-mirror/server/modules"
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
type createMessage struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Message []byte `json:"message"`
}

type moduleCreator modules.Module

//NewModuleContext creates a moduleContext with default set of modules,
//and module creators.
//The module context contains all modules initiated serverside and connects them
//to the right channels for sending and receiving messsages.
func NewModuleContext(writeChannel chan []byte, readChannel chan []byte) ModuleContext {
	var mods = make([]modules.Module, 0)
	//	mods = append(mods, systeminfo.NewSysInfoModule(writeChannel, "systeminfo", 200*time.Millisecond))

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
		var response createMessage
		err := json.Unmarshal(incoming, &response)
		if err != nil {
			continue
		}
		creator := m.Creators[response.Name]
		if creator == nil {
			continue
		}
		for _, mod := range m.Modules {
			//Prevent duplicate module initiations
			if mod.GetId() == response.ID {
				continue
			}
		}

		mod, err := creator.CreateFromMessage(response.Message, m.WriteChannel)
		if err == nil {
			m.Modules = append(m.Modules, mod)
			go mod.TimedUpdate()
		}

	}
}

//InitialMessages sends updates for all current modules
//Will be called when a new WS is established to send initial data.
func (m ModuleContext) InitialMessages() {
	for _, module := range m.Modules {
		module.Update()
	}

}
