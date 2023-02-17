package slackmodule

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/slack-go/slack"
)

func GetSlackProvider(apiKey string, channel string) SlackProvider {
	fmt.Println("Creating slack integration for channel: ", channel)
	return &SlackLiveProvider{
		api:       *slack.New(apiKey),
		channel:   channel,
		userNames: make(map[string]string),
	}
}

type SlackProvider interface {
	GetLatestMessages(noOfMessages int) []slack.Message
}

type SlackProviderError struct {
	err string
}

func (e SlackProviderError) Error() string {
	return e.err
}

func (p *SlackLiveProvider) GetLatestMessages(noOfMessages int) []slack.Message {

	if p.channelId == "" {
		channelId, err := FindChannelId(p.channel, p.api)
		if err == nil {
			p.channelId = *channelId
		}

		//emojis, err := GetEmojiList(p.api)
	}

	messages, _ := GetConversationHistory(p.channelId, p.api, noOfMessages)

	allMessages := make([]slack.Message, 0)
	for _, message := range messages {
		allMessages = append(allMessages, message)
		if message.ReplyCount > 0 {
			threadMessages, _ := GetConversationHistory(message.Timestamp, p.api, 10)
			allMessages = append(allMessages, threadMessages...)
		}
	}

	updatedMessages := p.addUserNames(allMessages)
	return updatedMessages

}

func (p *SlackLiveProvider) addUserNames(messages []slack.Message) []slack.Message {
	updatedMessages := make([]slack.Message, len(messages))
	for i := range messages {
		mess := messages[i]
		if len(mess.Username) > 0 {
			continue
		}

		name := p.getUserName(mess.Username)
		mess.Username = name
		p.replaceUserNameInText(&mess)
		updatedMessages[i] = mess

		fmt.Println(mess.Username)

	}
	return updatedMessages

}
func (p *SlackLiveProvider) getUserName(name string) string {
	cachedName, exists := p.userNames[name]
	if exists {
		return cachedName
	}
	user, err := GetUserName(name, p.api)
	if err != nil {
		return name
	}
	p.userNames[name] = user.RealName
	return user.RealName
}
func (p *SlackLiveProvider) replaceUserNameInText(mess *slack.Message) {
	re, _ := regexp.Compile("<@([A-Z0-9]+)>")

	matches := re.FindAllStringSubmatch(mess.Text, -1)

	replaceTarget := make(map[string]string)
	for main, submatches := range matches {
		for _, match := range submatches {
			user := p.getUserName(match)
			replaceTarget[submatches[main]] = user
		}
	}
	var returnText string = mess.Text
	for key, val := range replaceTarget {
		returnText = strings.ReplaceAll(returnText, key, val)
	}
	mess.Text = returnText
}
func GetUserName(userId string, client slack.Client) (*slack.User, error) {
	return client.GetUserInfo(userId)
}

func GetEmojiList(client slack.Client) (map[string]string, error) {
	return client.GetEmoji()
}
func GetConversationHistory(channelName string, client slack.Client, limit int) ([]slack.Message, error) {
	params := slack.GetConversationHistoryParameters{
		ChannelID: channelName,
		Limit:     limit,
	}

	response, err := client.GetConversationHistory(&params)
	if err != nil {
		fmt.Errorf("Error while fetching slack message, %s", err.Error())
		return make([]slack.Message, 0), err
	}
	return response.Messages, nil
}

func FindChannelId(channelName string, client slack.Client) (*string, error) {
	baseCursor := ""
	fmt.Println("Attempting to list slack channels")
	for {
		channels, nextCursor, err := client.GetConversations(&slack.GetConversationsParameters{Cursor: baseCursor})
		if err != nil {
			return nil, err
		}

		baseCursor = nextCursor
		for i, channel := range channels {
			fmt.Println(i, "Checked channel: ", channel.Name, " expecting ", channelName)
			if channel.Name == channelName {
				return &channel.ID, nil
			}
		}
		if nextCursor == "" {
			return nil, SlackProviderError{err: "Channel with id: " + channelName + "not found"}
		}
	}
}

type SlackLiveProvider struct {
	api       slack.Client
	channel   string
	channelId string
	userNames map[string]string
	emojis    map[string]string
}
