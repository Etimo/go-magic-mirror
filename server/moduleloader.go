package server

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/etimo/go-magic-mirror/server/modules"
	"github.com/etimo/go-magic-mirror/server/modules/systeminfo"
)

type ModuleContext struct {
	Modules      []modules.Module
	Creators     map[string]moduleCreator
	WriteChannel chan []byte
	ReadChannel  chan []byte
}

type moduleCreator modules.Module

type createMessage struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Message []byte `json:"message"`
}

func NewModuleContext(channel chan []byte) ModuleContext {
	var mods = make([]modules.Module, 0)
	mods = append(mods, systeminfo.NewSysInfoModule(channel, "systeminfo", 200*time.Millisecond))

	moduleCreator := map[string]moduleCreator{
		"systeminfo": systeminfo.SysinfoModule{},
	}
	return ModuleContext{
		Modules:      mods,
		Creators:     moduleCreator,
		WriteChannel: channel,
	}

}
func (m ModuleContext) SetupTimedUpdates() {

	for _, module := range m.Modules {
		fmt.Printf("module: %v\n", module)
		go module.TimedUpdate()
	}
}
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
			if mod.GetId() == response.Id {
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
func (m ModuleContext) InitialMessages() {
	for _, module := range m.Modules {
		module.Update()
	}

}
