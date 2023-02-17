package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/etimo/go-magic-mirror/server/modules/weather"

	"github.com/etimo/go-magic-mirror/server/modules"
	"github.com/etimo/go-magic-mirror/server/modules/clock"
	"github.com/etimo/go-magic-mirror/server/modules/googlecal"
	"github.com/etimo/go-magic-mirror/server/modules/photomod"
	slackmodule "github.com/etimo/go-magic-mirror/server/modules/slack"
	"github.com/etimo/go-magic-mirror/server/modules/systeminfo"
)

//NewModuleContext creates a moduleContext with default set of modules,
//and module creators.
//The module context contains all modules initiated serverside and connects them
//to the right channels for sending and receiving messsages.
func NewModuleContext(writeChannel chan []byte, readChannel chan []byte, callbackChannel chan bool) ModuleContext {
	var mods = make([]modules.Module, 0)
	var layouts = make(map[string]Layout, 0)
	var delay = time.Duration(1000)

	mods = append(mods, photomod.NewPhotoModule(writeChannel, "logo", "./public/etimo-logo-white.svg", delay*time.Millisecond))
	layouts["logo"] = Layout{X: 1, Y: 1, Width: 1, Height: 1, PluginType: "Photo"}
	mods = append(mods, clock.NewClockModule(writeChannel, "clock", 4, 4, 2, 2, delay*time.Millisecond))
	layouts["clock"] = Layout{X: 5, Y: 1, Width: 3, Height: 1, PluginType: "Clock"}
	mods = append(mods, photomod.NewPhotoModule(writeChannel, "photo", "https://homepages.cae.wisc.edu/~ece533/images/arctichare.png", 5*delay*time.Millisecond))
	layouts["photo"] = Layout{X: 1, Y: 3, Width: 2, Height: 2, PluginType: "Photo"}
	mods = append(mods, weather.NewWeatherModule(writeChannel, "weather", 1, 2, 2, 1, delay*15*time.Millisecond))
	layouts["weather"] = Layout{X: 4, Y: 3, Width: 1, Height: 1}
	mods = append(mods, weather.NewWeatherModule(writeChannel, "weather", 1, 2, 2, 1, delay*15*time.Millisecond))
	layouts["weather"] = Layout{X: 4, Y: 3, Width: 1, Height: 1}

	slackModule := slackmodule.NewSlackModule(
		writeChannel,
		"slackannounce",
		100*time.Second,
		100*time.Second,
		slackmodule.GetSlackProvider(
			os.Getenv("slackToken"),
			"etimo_internal"))
	mods = append(mods, slackModule)
	layouts[slackModule.Id] = Layout{X: 1, Y: 5, Width: 5, Height: 6}

	writer := json.NewEncoder(models.ChannelWriter{Channel: writeChannel})

	moduleCreator := map[string]moduleCreator{
		"systeminfo":     systeminfo.SysinfoModule{},
		"googlecalendar": googlecal.GoogleCalendarModule{},
	}

	return ModuleContext{
		Modules:         mods,
		Creators:        moduleCreator,
		WriteChannel:    writeChannel,
		ReadChannel:     readChannel,
		CallbackChannel: callbackChannel,
		Layouts:         layouts,
		MessageWriter:   writer,
	}

}

//RecieveCreateMessage handles incoming messages from the frontend and initiate
//modules on the server. This can be used instead of creating them on the server
//during construction. Message sent from frontend must match the createMessage
//struct and each module places own demands on the inner message.
func RecieveCreateMessage(m *ModuleContext) {
	for {
		incoming := <-m.ReadChannel
		var request CreateMessage

		err := json.Unmarshal(incoming, &request)
		if err != nil {
			continue
		}
		log.Printf("Received createione %v\n", request)
		handleMessage(request, incoming, m)
	}
}

func handleMessage(request CreateMessage, incoming []byte, m *ModuleContext) {

	creator := m.Creators[request.Name]
	if creator == nil {
		return
	}

	for _, mod := range m.Modules {
		//Prevent duplicate module initiations
		if mod.GetId() == request.Id {
			log.Printf("There is already a module with Id: %s", request.Id)
			return
		}
	}

	mod, err := creator.CreateFromMessage(incoming, m.WriteChannel)
	if err == nil && mod.GetId() == request.Id {
		m.Modules = append(m.Modules, mod)
		go mod.TimedUpdate()
		log.Printf("Added %v %v %d!", mod, err, len(m.Modules))
	}
}

func ReadCallback(m *ModuleContext) {
	for {
		<-m.CallbackChannel
		initialMessages(m)
	}
}

//InitialMessages sends updates for all current modules
//Will be called when a new WS is established to send initial data.
func initialMessages(m *ModuleContext) {
	for _, mod := range m.Modules {
		fmt.Printf("Updating module: %v", mod)
		mod.Update()
	}
	m.sendLayouts()
}
