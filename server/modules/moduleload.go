package modules

import (
	"time"

	"github.com/etimo/go-magic-mirror/server/modules/systeminfo"
)

type ModuleContext struct {
	Modules []Module
}

var one = systeminfo.SysMessage{}

func NewModuleContext(channel chan []byte) ModuleContext {
	var mods = []Module{
		systeminfo.NewSysInfoModule(channel,
			"systeminfo",
			200*time.Millisecond),
	}
	return ModuleContext{Modules: mods}

}
func (m ModuleContext) SetupTimedUpdates() {

	for _, module := range m.Modules {
		go module.TimedUpdate()
	}
}
func (m ModuleContext) InitialMessages() {
	for _, module := range m.Modules {
		module.Update()
	}

}
