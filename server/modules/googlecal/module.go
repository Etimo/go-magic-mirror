package googlecal

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/etimo/go-magic-mirror/server/modules"
)

const configEnvVariable = "MAGIC_MIRROR_SERVICE_LOCATION"

//EventSource interface represents a source of calendar events for the plugin.
type EventSource interface {
	GetEvents(startDateTime string,
		stopDateTime string,
		numberOfEvents int, initialLoog bool) []UpdateMessage
}

//CreateMessage represents the JSON structure sent from frontend to represent the creation of a new Calendar module.
type CreateMessage struct {
	Calendars []string `json:"calendars"`
	Id        string   `json:"Id"`
}

//GoogleCalendarModule is a calendar module for the mirror, built initially for use with Google Calendar.
type GoogleCalendarModule struct {
	APIKeyFile  string
	Calendars   []string
	Id          string
	Channel     chan []byte
	EventSource EventSource
}

//EventMessage represents the JSON structure sent to frontend for each event.
type EventMessage struct {
	Summary     string `json:"summary"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	ColorId     string `json:"colorId"`
	CreatorName string `json:"creatorName"`
}

//UpdateMessage represents the JSON structure for a calendar and it's containing events.
type UpdateMessage struct {
	CalendarName string         `json:"calendarName"`
	Events       []EventMessage `json:"currentEvents"`
}

//Message represents the Json structure sent to the frontend. It contains information on the plugin Id and calendar events.
type Message struct {
	Id      string          `json:"Id"`
	Type    string          `json:"type"`
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

//NewGoogleCalendarModuleAlternativeSource allows the use of the googleCalendar module with different event backends.
func NewGoogleCalendarModuleAlternativeSource(es EventSource,
	Id string,
	writeChannel chan []byte) *GoogleCalendarModule {
	return &GoogleCalendarModule{
		Id:          Id,
		EventSource: es,
		Channel:     writeChannel,
	}

}

//NewGoogleCalendarModule Initializes a new Google calendar module with the default GoogleCalendar calendar event source.
func NewGoogleCalendarModule(
	Id string,
	calendarDescriptors []string,
	writeChannel chan []byte,
) (*GoogleCalendarModule, error) {
	config := os.Getenv(configEnvVariable)
	eventSource, err := createGoogleCalendarSource(config, calendarDescriptors)
	if err != nil {
		log.Println("Could not create google calendar plugin with Id:,",
			Id, ": ", err.Error())
		return nil, err
	}
	return &GoogleCalendarModule{
		Id:          Id,
		EventSource: eventSource,
		Channel:     writeChannel,
	}, nil
}

/**
*initialLoad sets if the eventsource should check for NEW events
*before sending and update message.
 */
func (gc *GoogleCalendarModule) getEvents(initialLoad bool) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day()-2, 0, 0, 0, 0, time.Local)
	oneMoreWeek := startOfDay.Add(time.Hour * 24 * 100)
	events := gc.EventSource.GetEvents(
		startOfDay.Format(time.RFC3339),
		oneMoreWeek.Format(time.RFC3339),
		99,
		initialLoad)
	if events == nil {
		return
	}

	bytes, _ := json.Marshal(
		Message{
			Id:      gc.Id,
			Updates: events,
			Type:    "GoogleCalendar",
		},
	)
	gc.Channel <- bytes
}

//Update triggers a push of new information from module to linked channels.
func (gc GoogleCalendarModule) Update() {
	gc.getEvents(true)
}

//TimedUpdate triggers an update and a sleep to allow for delayed periodic updates.
func (gc GoogleCalendarModule) TimedUpdate() {
	for {
		gc.getEvents(true)
		time.Sleep(time.Second * 15)
	}
}

//GetId returns the ID of the module.
func (gc GoogleCalendarModule) GetId() string {
	return gc.Id
}

//CreateFromMessage attempts to unmarshal the provided bytes into a google calendar module creation message, if succesfull the information will be used to initialize a new module.
func (gc GoogleCalendarModule) CreateFromMessage(messageBytes []byte, channel chan []byte) (modules.Module, error) {
	var module GoogleCalendarModule
	var message CreateMessage
	err := json.Unmarshal(messageBytes, &message)
	if err != nil {
		log.Println("Could not construct google calendar message from bytes")
		return module, err
	}
	return NewGoogleCalendarModule(message.Id, message.Calendars, channel)

}
