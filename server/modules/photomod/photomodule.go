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
	delay  time.Duration
}

type PhotoMessage struct {
	Id   string `json:"Id"`
	Type string `json:"type"`
	Url  string `json:"Url"`
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
			Id:   c.Id,
			Type: "Photo",
			Url:  c.Url,
		}
	} else {

		message = PhotoMessage{
			Id:   c.Id,
			Type: "Photo",
			Url:  c.GetRandomPhoto(),
		}
	}

	c.writer.Encode(message)
}

func NewPhotoModule(
	channel chan []byte,
	Id string,
	Url string,
	delayInfoPush time.Duration) PhotoModule {
	return PhotoModule{
		writer: json.NewEncoder(models.ChannelWriter{Channel: channel}),
		Id:     Id,
		Url:    Url,
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
