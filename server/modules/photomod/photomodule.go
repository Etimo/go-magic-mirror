package photomod

import (
	"encoding/json"
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
	Id  string `json:"Id"`
	Url string `json:"Url"`
}

// structure function  PhotoModule.Update();
func (c PhotoModule) Update() {

	c.writer.Encode(PhotoMessage{
		Id:  c.Id,
		Url: c.Url,
	})
}

func NewPhotoModule(channel chan []byte,
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
