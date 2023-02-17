package support

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/slack-go/slack"
)

var fakeSlackSourceMessage slack.Message = slack.Message{
	Msg: slack.Msg{
		Username:  "ERIK",
		Timestamp: "1234",
		Text:      "JAG ÄR SÅ BALL",
	},
}

type FakeSlackProvider struct{}

func (FakeSlackProvider) GetLatestMessages(max int) []slack.Message {
	return []slack.Message{
		fakeSlackSourceMessage,
	}
}

func CodepointToUnicode(word string) string {

	var ret string

	for _, chunk := range strings.Split(strings.TrimPrefix(strings.ToUpper(word), "U+"), "U+") {
		c, err := strconv.ParseInt(chunk, 16, 64)
		if err != nil {
			return ret
		}
		ret = fmt.Sprintf("%s%c", ret, c)
	}
	return ret
}

func TestEmojiSource(t *testing.T) {
	fmt.Println("Emoji source test")
	//Capture stdout
	var buf bytes.Buffer
	log.SetOutput(&buf)
	str, _ := hex.DecodeString("0023-FE0F-20E3")
	fmt.Println(string(str))
	fmt.Printf("0023-FE0F-20E3")

	fmt.Println("Stdout: ", buf.String())

	emojiSource := GetEmojiSource()
	testEmoji := ":slightly_smiling_face:"
	emojiSource.ReplaceEmojiInString(testEmoji)
	fmt.Println("emojiSource: ", emojiSource)
	if emojiSource == nil {
		fmt.Println("Couldn't get file")

		t.Fail()
	}
}

func TestReplaceEmojis(t *testing.T) {
	emojis := EmojiSource{
		Emojis: []EmojiEntry{
			EmojiEntry{
				ShortName: "fake-emoji",
				Unified:   "a",
			},
			EmojiEntry{ShortName: "a-better-fake", Unified: "b"},
		},
		MapMoji: make(map[string]EmojiEntry),
	}
	for _, e := range emojis.Emojis {
		emojis.MapMoji[e.ShortName] = e
	}
	emojis.init()
	baseString := ":fake-emoji: hello! :a-better-fake:"
	expectedString := "a hello! b"
	withEmojis := emojis.ReplaceEmojiInString(baseString)
	if withEmojis != expectedString {
		t.Fail()
	}

}

func TestFindBasePath(t *testing.T) {
	path, _ := AttemptToFindBasePath()
	if strings.Contains(path, ".") {
		t.Fail()

	}

}
