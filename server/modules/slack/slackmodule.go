package slackmodule

import (
	"encoding/json"
	"time"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/etimo/go-magic-mirror/server/models/support"
	"github.com/etimo/go-magic-mirror/server/modules"
	"github.com/slack-go/slack"
)

const slackAPITokentEnv string = "slackApiToken"

var emojis *support.EmojiSource = support.GetEmojiSource()

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
	api         *SlackProvider
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
		api:         &provider,
	}
}

func replaceEmojis(messages []slack.Message) []slack.Message {

	for i := range messages {
		messages[i].Text = emojis.ReplaceEmojiInString(messages[i].Text)
		//fmt.Println(messages[i].Text)
	}
	return messages

}

func replaceUsenames(text string, usermap map[string]string) {

}
func (c SlackModule) Update() {

	client := *c.api
	messages := client.GetLatestMessages(5)
	messages = replaceEmojis(messages)

	slackUpdateMessage := SlackUpdateMessage{
		Id:            c.Id,
		SlackMessages: messages,
		Type:          "Slack",
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

func (c SlackModule) CreateFromMessage(
	message []byte, channel chan []byte) (modules.Module, error) {
	return SlackModule{}, nil
}
