package clock

import (
	"encoding/json"
	"time"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/etimo/go-magic-mirror/server/modules"
)

type ClockMessage struct {
	ID   string    `json:"id"`
	Date ClockDate `json:"date"`
	Time ClockTime `json:"time"`
}
type ClockDate struct {
	Day   int    `json:"day"`
	Month string `json:"month"`
	Year  int    `json:"year"`
}
type ClockTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}
type CreateMessage struct {
	ID    string `json:"id"`
	Delay int    `json:"delay"`
}
type ClockModule struct {
	writer *json.Encoder
	id     string
	delay  time.Duration
}

func NewClockModule(channel chan []byte,
	id string,
	delayInfoPush time.Duration) ClockModule {
	return ClockModule{
		writer: json.NewEncoder(models.ChannelWriter{Channel: channel}),
		id:     id,
		delay:  delayInfoPush,
	}
}

func (c ClockModule) Update() {
	var message ClockMessage

	var timeNow = time.Now()

	message.ID = c.GetId()
	message.Time.Hour = timeNow.Hour()
	message.Time.Minute = timeNow.Minute()
	message.Time.Second = timeNow.Second()
	message.Date.Day = timeNow.Day()
	message.Date.Month = timeNow.Month().String()
	message.Date.Year = timeNow.Year()
	c.writer.Encode(message)
}

func (c ClockModule) GetId() string {
	return c.id
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
	return NewClockModule(channel, targetMessage.ID, time.Duration(targetMessage.Delay)*time.Millisecond), nil
}
