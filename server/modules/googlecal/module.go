package googlecal

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/etimo/go-magic-mirror/server/modules"
)

const CONFIG_ENV_VARIABLE = "MAGIC_MIRROR_SERVICE_LOCATION"

type EventSource interface {
	GetEvents(startDateTime string,
		stopDateTime string,
		numberOfEvents int, initialLoog bool) []UpdateMessage
}

type CreateMessage struct {
	Calendars []string `json:"calendars"`
	Id        string   `json:"id"`
}
type GoogleCalendarModule struct {
	ApiKeyFile  string
	Calendars   []string
	Id          string
	Channel     chan []byte
	EventSource EventSource
}
type EventMessage struct {
	Summary     string `json:"summary"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	ColorId     string `json:"colorId"`
	CreatorName string `json:"creatorName"`
}
type UpdateMessage struct {
	CalendarName string         `json:"calendarName"`
	Events       []EventMessage `json:"currentEvents"`
}
type Message struct {
	Id      string          `json:"calendarName"`
	Updates []UpdateMessage `json:"calendars"`
}

func matchCalendar(expressions []string,
	calendarString string) bool {
	for _, val := range expressions {
		match, _ := regexp.MatchString(val, calendarString)
		if match {
			return true
		}

	}
	return false
}
func NewGoogleCalendarModuleAlternativeSource(es EventSource,
	id string,
	writeChannel chan []byte) GoogleCalendarModule {
	return GoogleCalendarModule{
		Id:          id,
		EventSource: es,
		Channel:     writeChannel,
	}

}
func NewGoogleCalendarModule(
	id string,
	calendarDescriptors []string,
	writeChannel chan []byte,
) GoogleCalendarModule {
	config := os.Getenv(CONFIG_ENV_VARIABLE)
	eventSource, err := createGoogleCalendarSource(config, calendarDescriptors)
	if err != nil {
		log.Println("Could not create google calendar plugin with ID:,",
			id, ": ", err.Error())
	}
	return GoogleCalendarModule{
		Id:          id,
		EventSource: eventSource,
		Channel:     writeChannel,
	}
}

/**
*initialLoad sets if the eventsource should check for NEW events
*before sending and update message.
 */
func (gc GoogleCalendarModule) getEvents(initialLoad bool) {
	now := time.Now() //.Format(time.RFC3339)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	oneMoreWeek := startOfDay.Add(time.Hour * 24 * 7)
	events := gc.EventSource.GetEvents(
		startOfDay.Format(time.RFC3339),
		oneMoreWeek.Format(time.RFC3339),
		99,
		initialLoad)
	bytes, _ := json.Marshal(
		Message{
			Id:      gc.Id,
			Updates: events,
		},
	)
	gc.Channel <- bytes
}
func (gc GoogleCalendarModule) Update() {
	gc.getEvents(true)
}
func (gc GoogleCalendarModule) TimedUpdate() {
	gc.getEvents(false)
	time.Sleep(time.Second * 5)
}
func (gc GoogleCalendarModule) GetId() string {
	return gc.Id
}
func (gc GoogleCalendarModule) CreateFromMessage(messageBytes []byte, channel chan []byte) (modules.Module, error) {
	var module GoogleCalendarModule
	var message CreateMessage
	err := json.Unmarshal(messageBytes, &message)
	if err != nil {
		log.Println("Could not construct google calendar message from bytes")
		return module, err
	}
	return NewGoogleCalendarModule(message.Id, message.Calendars, channel), nil

}
