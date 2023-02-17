package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type EmojiSource struct {
	Emojis      []EmojiEntry
	MapMoji     map[string]EmojiEntry
	MatchRegexp *regexp.Regexp
}

/*
{"name":"HASH KEY",
"unified":"0023-FE0F-20E3",
"non_qualified":"0023-20E3","docomo":"E6E0","au":"EB84",
"softbank":"E210",
"google":"FE82C",
"image":"0023-fe0f-20e3.png",
"sheet_x":0,"sheet_y":0,"short_name":"hash",
"short_names":["hash"],"text":null,
"texts":null,
"category":"Symbols",
"subcategory":"keycap",
"sort_order":1500,
"added_in":"0.6",
"has_img_apple":true,
"has_img_google":true,
"has_img_twitter":true,
"has_img_facebook":false}
*/
type EmojiEntry struct {
	Name      string `json:"name"`
	Unified   string `json:"unified"`
	ShortName string `json:"short_name"`
}

func GetEmojiSource() *EmojiSource {
	file := getEmojiFile()
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("Error while constructing emoji-support %s\n", err.Error())
		return nil
	}

	emojiSource := EmojiSource{}
	emojiSource.MapMoji = make(map[string]EmojiEntry)

	emojiSource.init()

	errAgain := json.Unmarshal(bytes, &emojiSource.Emojis)

	for _, entry := range emojiSource.Emojis {
		entry.Unified = unifiedToString(entry.Unified)
		emojiSource.MapMoji[entry.ShortName] = entry
	}
	if errAgain != nil {
		fmt.Printf("Error while constructing emoji-support %s\n", err.Error())
		return nil
	}

	return &emojiSource
}
func unifiedToString(unified string) string {
	c, err := strconv.ParseInt(unified, 16, 64)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%c", c)
}
func (s *EmojiSource) init() {

	regexp, _ := regexp.Compile(":([a-z-_]+):")
	s.MatchRegexp = regexp

}
func (s EmojiSource) ReplaceEmojiInString(text string) string {
	matches := s.MatchRegexp.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return text
	}
	replaceTarget := make(map[string]string)
	for _, match := range matches {
		val, exists := s.MapMoji[match[1]]
		if exists {
			replaceTarget[match[0]] = val.Unified
		}
	}
	var returnText string = text
	for key, val := range replaceTarget {
		returnText = strings.ReplaceAll(returnText, key, val)
	}
	return returnText
}
func getEmojiFile() string {
	source, found := os.LookupEnv("emojijsonsource")
	var finalSource string

	if !found {
		path, _ := AttemptToFindBasePath()
		finalSource = path + "/resources/emoji.json"
	} else {
		finalSource = source
	}

	fmt.Println("Using emoji source: ", finalSource)
	return finalSource
}
