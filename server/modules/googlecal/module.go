package googlecal

import (
	"encoding/json"
	"log"
	"regexp"
	"time"
)

type EventSource interface {
	GetEvents(startDateTime string,
		stopDateTime string,
		numberOfEvents int) []UpdateMessage
}

type createMessage struct {
	ApiKeyFile string   `json:"keyFile"`
	Calendars  []string `json:"calendars"`
	Id         string   `json:"id"`
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
	credentialLocation string,
	id string,
	calendarDescriptors []string,
	writeChannel chan []byte,
) GoogleCalendarModule {
	eventSource, err := createGoogleCalendarSource(credentialLocation, calendarDescriptors)
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

func (gc GoogleCalendarModule) Update() {
	now := time.Now() //.Format(time.RFC3339)
	startOfDay := time.Date(now.Year(), now.Month(), now.Year(), 0, 0, 0, 0, time.Local)
	oneMoreWeek := startOfDay.Add(time.Hour * 24 * 7)
	events := gc.EventSource.GetEvents(startOfDay.Format(time.RFC3339),
		oneMoreWeek.Format(time.RFC3339),
		99)
	bytes, _ := json.Marshal(events)
	gc.Channel <- bytes
}
func (gc GoogleCalendarModule) TimedUpdate() {
	time.Sleep(time.Second * 15)
	gc.Update()
}
func (gc GoogleCalendarModule) GetId() string {
	return gc.Id
}
