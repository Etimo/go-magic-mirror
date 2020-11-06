package photomod

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/etimo/go-magic-mirror/server/models"
)

type PhotoModule struct {
	writer *json.Encoder
	Id     string
	Url    string
	Height string
	Width  string
	delay  time.Duration
}

type PhotoMessage struct {
	Id     string `json:"Id"`
	Type   string `json:"type"`
	Url    string `json:"Url"`
	Height string `json:"Height"`
	Width  string `json:"Width"`
}

func (c PhotoModule) GetRandomPhoto() string {
	//var x = []string
	files, err := ioutil.ReadDir("./public/photos")
	if err != nil {
		log.Fatal(err)
	}

	return "./public/photos/" + files[rand.Intn(len(files))].Name()
}

// structure function  PhotoModule.Update();
func (c PhotoModule) Update() {
	var message PhotoMessage

	if c.Id == "logo" {
		message = PhotoMessage{
			Id:     c.Id,
			Type:   "Photo",
			Url:    c.Url,
			Height: c.Height,
			Width:  c.Width,
		}
	} else {

		message = PhotoMessage{
			Id:     c.Id,
			Type:   "Photo",
			Url:    c.GetRandomPhoto(),
			Height: c.Height,
			Width:  c.Width,
		}
	}

	c.writer.Encode(message)
}

func NewPhotoModule(channel chan []byte,
	Id string,
	Url string,
	Height string,
	width string,
	delayInfoPush time.Duration) PhotoModule {
	return PhotoModule{
		writer: json.NewEncoder(models.ChannelWriter{Channel: channel}),
		Id:     Id,
		Url:    Url,
		Height: Height,
		Width:  width,
		delay:  delayInfoPush,
	}
}

// structure function  PhotoModule.TimedUpdate();
func (c PhotoModule) TimedUpdate() {
	for {
		time.Sleep(c.delay)
		c.Update()
	}
}

// structure function  PhotoModule.GetId();
func (c PhotoModule) GetId() string {
	return c.Id
}
