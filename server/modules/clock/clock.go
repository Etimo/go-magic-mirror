package clock

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/etimo/go-magic-mirror/server/modules"
)

type ClockMessage struct {
	Id   string    `json:"Id"`
	Type string    `json:"type"`
	Date ClockDate `json:"date"`
	Time ClockTime `json:"time"`
}
type ClockDate struct {
	Day   int    `json:"day"`
	Month string `json:"month"`
	Year  int    `json:"year"`
}
type ClockTime struct {
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
	Second string `json:"second"`
}
type CreateMessage struct {
	Id    string `json:"Id"`
	Delay int    `json:"delay"`
}
type ClockModule struct {
	writer *json.Encoder
	Id     string
	delay  time.Duration
}

func NewClockModule(channel chan []byte,
	Id string,
	delayInfoPush time.Duration) ClockModule {
	return ClockModule{
		writer: json.NewEncoder(models.ChannelWriter{Channel: channel}),
		Id:     Id,
		delay:  delayInfoPush,
	}
}

func (c ClockModule) Update() {
	var message ClockMessage

	var timeNow = time.Now()

	message.Id = c.GetId()
	message.Time.Hour = FormatTime(timeNow.Hour())
	message.Time.Minute = FormatTime(timeNow.Minute())
	message.Time.Second = FormatTime(timeNow.Second())
	message.Date.Day = timeNow.Day()
	message.Date.Month = timeNow.Month().String()
	message.Date.Year = timeNow.Year()
	message.Type = "Clock"
	c.writer.Encode(message)
}

func FormatTime(time int) string {
	formattedTime := strconv.Itoa(time)
	if len(formattedTime) < 2 {
		formattedTime = "0" + formattedTime
	}
	return formattedTime
}
func (c ClockModule) GetId() string {
	return c.Id
}

func (c ClockModule) TimedUpdate() {
	for {
		time.Sleep(c.delay)
		c.Update()
	}
}

func (c ClockModule) CreateFromMessage(message []byte, channel chan []byte) (modules.Module, error) {
	var targetMessage CreateMessage
	err := json.Unmarshal(message, &targetMessage)
	if err != nil {
		return nil, err
	}
	return NewClockModule(channel, targetMessage.Id, time.Duration(targetMessage.Delay)*time.Millisecond), nil
}
