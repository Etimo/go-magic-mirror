package googlecal

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type GoogleCalendarModule struct {
	ApiKeyFile string
	Calendars  []string
	Id         string
	Channel    chan []byte
}

/*func connectWithToken(config *oauth2.Config, tokenFile string) (*http.Client, error) {
	token, err := readToken(tokenFile, config)
	if err != nil {
		token = queryWebToken(config)
	}
	return config.Client(context.Background(), token), nil
}

// Use a JSON token for a service account, this service account should have relevant calendars shared.
func readToken(file string, config *oauth2.Config) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		queryWebToken(config)
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok) //If it's not great JSON this won't work.
	return tok, err
}

// Server auth flow.
func queryWebToken(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Link calendar at this URL:\n%v\n", authURL)
	fmt.Println("Token will be stored on disk. Do not accidentally commit.")

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to retrieve authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)

	if err != nil {
		log.Fatalf("Problem during token exchange: %v", err)
	}
	return tok
}*/

func NewGoogleCalendarModule(credentialLocation string) {
	f, ferr := ioutil.ReadFile(credentialLocation)
	if ferr != nil {
		log.Fatal("Can not start google calendar plugin: ", ferr.Error())
	}
	config, err := google.JWTConfigFromJSON(f, calendar.CalendarReadonlyScope)
	client := config.Client(oauth2.NoContext)
	cal, err := calendar.New(client)
	fmt.Printf("Cal: %v and %v\n", cal, err)
	list, _ := cal.CalendarList.List().Do()
	//fmt.Printf("List: %v and %v\n", list, errCal)
	fmt.Println("More cal: ", len(list.Items))
	for _, cal := range list.Items {
		fmt.Println("This is this calendar: ", cal.Id, cal.Summary)
	}
	cal.Events.List("Kurser").
}
func (gc GoogleCalendarModule) init() {
}
func (gc GoogleCalendarModule) Update() {
}
func (gc GoogleCalendarModule) TimedUpdate() {
}
func (gc GoogleCalendarModule) GetId() string {
	return gc.Id
}
