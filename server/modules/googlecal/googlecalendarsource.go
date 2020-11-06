package googlecal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type googleCalendarSource struct {
	configFile        string
	client            *calendar.Service
	calendars         []string
	googleCalendarIds []string
	syncTokens        map[string]*string
}

/**
* Calls google calendar API with stored syncToken, this returns only
* a list of events changed since last call.
 */
func (gc googleCalendarSource) CheckUpdated(calendarId string) bool {
	syncToken := gc.syncTokens[calendarId]
	if nil == syncToken {
		return true
	}
	newEvents, err := gc.client.Events.List(calendarId).SyncToken(*syncToken).MaxResults(1).
		OrderBy("updated").Do()
	if err != nil {
		return false
	}
	if len(newEvents.Items) > 0 {
		return true
	}
	return false
}

//GetEvents queries google calendar for a specific number of events situated within the enclosing period.
func (gc googleCalendarSource) GetEvents(
	startDateTime string, stopDateTime string, numberOfEvents int, initialLoad bool) []UpdateMessage {
	fmt.Println("Time: ", startDateTime, stopDateTime)
	returnMessages := make([]UpdateMessage, len(gc.googleCalendarIds))

	for i, calendarId := range gc.googleCalendarIds {
		if !initialLoad && !gc.CheckUpdated(calendarId) {
			log.Println("No updates for calendar: ", calendarId, " : ", gc.calendars[i])
			continue
		}
		eventMessages := gc.getEventMessages(startDateTime,
			stopDateTime,
			gc.calendars[i],
			calendarId)
		returnMessages[i] = UpdateMessage{
			CalendarName: gc.calendars[i],
			Events:       eventMessages,
		}
	}
	return returnMessages
}

func (gc googleCalendarSource) getEventMessages(startDateTime, stopDateTime, calendarName, calendarID string) []EventMessage {
	fmt.Println("EventTime: ", startDateTime, stopDateTime)
	list, err := gc.client.Events.List(calendarID).
		TimeMin(startDateTime).
		TimeMax(stopDateTime).
		Do()
	if err != nil {
		log.Println("Problem with calendar: ", calendarID, " : ", calendarName)
	}
	fmt.Println("Found: ", len(list.Items))
	eventMessages := make([]EventMessage, len(list.Items))
	for i, event := range list.Items {
		eventMessages[i] = createEventMessage(event)
		time.Now()
	}
	gc.syncTokens[calendarID] = &list.NextSyncToken
	return eventMessages
}

func createEventMessage(event *calendar.Event) EventMessage {
	return EventMessage{
		Summary:     event.Summary,
		StartTime:   event.Start.DateTime,
		EndTime:     event.End.DateTime,
		ColorId:     event.ColorId,
		CreatorName: event.Creator.Email,
	}
}
func createGoogleCalendarSource(
	credentialLocation string,
	calendarDescriptors []string,
) (EventSource, error) {
	f, ferr := ioutil.ReadFile(credentialLocation)
	if ferr != nil {
		log.Println("Can not start google calendar plugin: ", ferr.Error())
		return nil, errors.New("could not read credential file")
	}
	config, err := google.JWTConfigFromJSON(f, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Println("Error creating config from JWT: ", err.Error())
		return nil, errors.New("could not read the provided credential as JWT key")
	}
	client := config.Client(oauth2.NoContext)
	gocal, err := calendar.New(client)
	if err != nil {
		log.Println("Error creating calendar client: ", err.Error())
		return nil, errors.New("could not create google connection client")
	}
	//fmt.Printf("Cal: %v and %v\n", gocal, err)
	list, errCal := gocal.CalendarList.List().MaxResults(999).Do()
	if errCal != nil {
		log.Println("Could not list google calendars", errCal.Error())
		return nil, errors.New("could not list google calendars")
	}
	//fmt.Printf("List: %v and %v\n", list, errCal)
	//fmt.Println("More cal: ", len(list.Items))
	calendarNames := make([]string, 0)
	calendarIds := make([]string, 0)
	for _, cal := range list.Items {
		log.Println("Calendar: ", cal.Summary)
		if matchCalendar(calendarDescriptors, cal.Summary) || matchCalendar(calendarDescriptors, cal.Id) {
			calendarNames = append(calendarNames, cal.Summary)
			calendarIds = append(calendarIds, cal.Id)
		}
	}
	return googleCalendarSource{
		configFile:        credentialLocation,
		client:            gocal,
		googleCalendarIds: calendarIds,
		calendars:         calendarNames,
		syncTokens:        make(map[string]*string),
	}, nil
}
