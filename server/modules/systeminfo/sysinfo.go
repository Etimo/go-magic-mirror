package systeminfo

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type MemoryUnit struct {
	size uint64
	text string
}

var (
	KB MemoryUnit = MemoryUnit{size: 1000, text: "KB"}
	MB MemoryUnit = MemoryUnit{size: KB.size * 1000, text: "MB"}
	GB            = MemoryUnit{size: 1000 * MB.size, text: "GB"}
)

type SysMessage struct {
	Id                string  `json:"id"`
	Os                string  `json:"os"`
	HostName          string  `json:"hostName"`
	TotalMemory       string  `json:"memoryTotal"`
	UsedMemoryPercent float64 `json:"memoryUsedPercent"`
	MemoryUsed        string  `json:"usedMemory"`
	Cpus              []Cpu   `json:"cpus"`
	Uptime            uint64  `json:"uptime"`
}
type Cpu struct {
	ModelName   string
	Mhz         int
	Utilization float64
}
type SysinfoModule struct {
	writer          *json.Encoder
	id              string
	delay           time.Duration
	constantMessage SysMessage
}

func NewSysInfoModule(channel chan []byte,
	id string,
	delayInfoPush time.Duration) SysinfoModule {
	return SysinfoModule{
		writer:          json.NewEncoder(models.ChannelWriter{Channel: channel}),
		id:              id,
		delay:           delayInfoPush,
		constantMessage: getConstantInfo(),
	}

}

//Retrieve all unchanging info about the system.
func getConstantInfo() SysMessage {
	message := SysMessage{}
	info, errInfo := host.Info()
	if errInfo == nil {
		message.HostName = info.Hostname
		message.Os = info.OS
	}
	infoCores, errCpu := cpu.Info()
	if errCpu == nil {
		cpus := make([]Cpu, len(infoCores))
		for i, core := range infoCores {
			cpus[i] = Cpu{
				ModelName: core.ModelName,
				Mhz:       int(core.Mhz),
			}
		}
		message.Cpus = cpus
	}
	return message
}

func (s SysinfoModule) Update() {

	message := s.constantMessage
	message.Id = s.GetId()
	memReport, errMem := mem.VirtualMemory()
	if errMem == nil {
		message.TotalMemory = convertMemUnit(memReport.Total, GB)
		message.UsedMemoryPercent = math.Round(memReport.UsedPercent)
	}
	times, errTimes := cpu.Percent(0, true)
	if errTimes == nil {
		for i, util := range times {
			message.Cpus[i].Utilization = util
		}
	}
	uptime, err := host.Uptime()
	if err == nil {
		message.Uptime = uptime
	}
	s.writer.Encode(message)
}

func convertMemUnit(memoryBytes uint64, unit MemoryUnit) string {
	return fmt.Sprintf("%d %s", uint64(memoryBytes)/unit.size, unit.text)
}

func (s SysinfoModule) GetId() string {
	return s.id
}
func (s SysinfoModule) TimedUpdate() {
	for {
		time.Sleep(s.delay)
		s.Update()
	}
}
