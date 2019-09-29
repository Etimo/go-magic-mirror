package sysinfo

import (
	"time"
)

type sysinfo struct {
	channel chan []byte
	id      string
	delay   int
}

func (s sysinfo) update() {

}
func (s sysinfo) getId() string {
	return s.id
}
func (s sysinfo) timedUpdate() {
	time.Sleep(500 * time.Millisecond)
}
