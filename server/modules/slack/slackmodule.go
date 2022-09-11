package slackmodule

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/etimo/go-magic-mirror/server/modules"
	"github.com/slack-go/slack"
)

const slackAPITokentEnv string = "slackApiToken"

type SlackUpdateMessage struct {
	Id            string          `json:"Id"`
	Type          string          `json:"type"`
	SlackMessages []slack.Message `json:"slackMessages"`
}

type SlackModule struct {
	writer      *json.Encoder
	Id          string
	delay       time.Duration
	refreshTime time.Duration
	api         SlackProvider
}

func NewSlackModule(
	channel chan []byte,
	Id string,
	delayInfoPush time.Duration,
	refreshTime time.Duration,
	provider SlackProvider,
) SlackModule {
	return SlackModule{
		writer:      json.NewEncoder(models.ChannelWriter{Channel: channel}),
		Id:          Id,
		delay:       delayInfoPush,
		refreshTime: refreshTime,
		api:         provider,
	}
}

func (c SlackModule) Update() {
	fmt.Println("UPDATE!")
	messages := c.api.GetLatestMessages(10)
	slackUpdateMessage := SlackUpdateMessage{
		Id:            c.Id,
		SlackMessages: messages,
		Type:          "SlackUpdate",
	}
	c.writer.Encode(slackUpdateMessage)
}

func (c SlackModule) GetId() string {
	return c.Id
}

func (c SlackModule) TimedUpdate() {
	for {
		time.Sleep(c.delay)
		c.Update()
	}
}

func (c SlackModule) CreateFromMessage(message []byte, channel chan []byte) (modules.Module, error) {
	return SlackModule{}, nil
}
