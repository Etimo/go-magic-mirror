package server

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/etimo/go-magic-mirror/server/modules"
)

type Layout struct {
	X          int32
	Y          int32
	Height     int32
	Width      int32
	PluginType string
}

type LayoutMessage struct {
	Type       string `json:"type"`
	Id         string `json:"pluginId"`
	X          int32  `json:"x"`
	Y          int32  `json:"y"`
	Height     int32  `json:"height"`
	Width      int32  `json:"width"`
	PluginType string `json:"pluginType"`
}

//ModuleContext : struct that contains array of all server side modules,
//and connects them with the correct channel for sending messages.
type ModuleContext struct {
	Modules         []modules.Module
	Layouts         map[string]Layout
	Creators        map[string]moduleCreator
	WriteChannel    chan []byte
	ReadChannel     chan []byte
	CallbackChannel chan bool
	MessageWriter   *json.Encoder
}

//SetupTimedUpdates starts the timedUpdate flow for all modules in the module list
//Each plugin will start a goRoutine that pushes messages to the shared channel, where websocket
//Should only be called once on startup.
func (m ModuleContext) SetupTimedUpdates() {

	for _, module := range m.Modules {
		fmt.Printf("module: %v\n", module)
		go module.TimedUpdate()
	}
}
func (m ModuleContext) timedLayout() {
	for {
		time.Sleep(5 * time.Second)
		m.sendLayouts()
	}
}

func (m ModuleContext) sendLayouts() {
	for name, layout := range m.Layouts {
		var layoutMessage LayoutMessage = LayoutMessage{
			Id:         name,
			X:          layout.X,
			Y:          layout.Y,
			Height:     layout.Height,
			Width:      layout.Width,
			Type:       "layout",
			PluginType: layout.PluginType,
		}
		fmt.Printf("%v\n", layoutMessage)
		m.MessageWriter.Encode(layoutMessage)
	}

}

//Todo: Break out into separate file in package
type CreateMessage struct {
	Name string `json:"name"`
	Id   string `json:"Id"`
}

type moduleCreator interface {
	CreateFromMessage([]byte, chan []byte) (modules.Module, error)
}
