package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/etimo/go-magic-mirror/server/models"
	"github.com/etimo/go-magic-mirror/server/modules"
)

type CreateMessage struct {
	Id    string `json:"Id"`
	Delay int    `json:"delay"`
}
type SMHIResponse struct {
	ApprovedTime  string      `json:approvedTime`
	ReferenceTime string      `json:referenceTime`
	TimeSeries    []TimeSerie `json:timeSeries`
}
type TimeSerie struct {
	ValidTime  string      `json:validTime`
	Parameters []Parameter `json:parameters`
}
type Parameter struct {
	Name      string    `json:name`
	LevelType string    `json:levelType`
	Level     int       `json:level`
	Unit      string    `json:unit`
	Values    []float64 `json:values`
}

var weatherTypes = map[uint8]string{
	1:  "wi wi-darksky-clear-day",
	2:  "wi wi-wu-mostlysunny",
	3:  "VariableCloudiness",
	4:  "HalfclearSky",
	5:  "CloudySky",
	6:  "Overcast",
	7:  "Fog",
	8:  "LightRainShowers",
	9:  "ModerateRainShowers",
	10: "HeavyRainShowers",
	11: "Thunderstorm",
	12: "LightSleetShowers",
	13: "ModerateSleetShowers",
	14: "HeavySleetShowers",
	15: "LightSnowShowers",
	16: "ModerateSnowShowers",
	17: "HeavySnowShowers",
	18: "LightRain",
	19: "ModerateRain",
	20: "HeavyRain",
	21: "Thunder",
	22: "LightSleet",
	23: "ModerateSleet",
	24: "HeavySleet",
	25: "LightSnowfall",
	26: "ModerateSnowfall",
	27: "HeavySnowfall",
}

type WeatherModule struct {
	writer *json.Encoder
	Id     string
	delay  time.Duration
}

func NewWeatherModule(channel chan []byte,
	Id string,
	delayInfoPush time.Duration) WeatherModule {
	return WeatherModule{
		writer: json.NewEncoder(models.ChannelWriter{Channel: channel}),
		Id:     Id,
		delay:  delayInfoPush,
	}
}

func (c WeatherModule) Update() {
	var weather = GetWeather("18.06", "59.33")
	var temperature = strconv.FormatFloat(weather.TimeSeries[0].Parameters[1].Values[0], 'f', -1, 64) + " " + string(weather.TimeSeries[0].Parameters[1].Unit)
	var weatherType = weatherTypes[uint8(weather.TimeSeries[0].Parameters[18].Values[0])]
	var message models.TextWidget
	message.Init(c.GetId(), 1, 1, temperature)
	message.Icon = weatherType
	c.writer.Encode(message)
}

func FormatTime(time int) string {
	formattedTime := strconv.Itoa(time)
	if len(formattedTime) < 2 {
		formattedTime = "0" + formattedTime
	}
	return formattedTime
}
func (c WeatherModule) GetId() string {
	return c.Id
}

func (c WeatherModule) TimedUpdate() {
	for {
		time.Sleep(c.delay)
		c.Update()
	}
}

func GetWeather(long string, lat string) *SMHIResponse {

	resp, err := http.Get(fmt.Sprintf("https://opendata-download-metfcst.smhi.se/api/category/pmp3g/version/2/geotype/point/lon/%v/lat/%v/data.json", long, lat))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		panic(err)
	}

	var response = new(SMHIResponse)
	json.Unmarshal(body, response)
	return response
}

func (c WeatherModule) CreateFromMessage(message []byte, channel chan []byte) (modules.Module, error) {
	var targetMessage CreateMessage
	err := json.Unmarshal(message, &targetMessage)
	if err != nil {
		return nil, err
	}
	return NewWeatherModule(channel, targetMessage.Id, time.Duration(targetMessage.Delay)*time.Millisecond), nil
}
