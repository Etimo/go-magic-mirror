package models

type ChannelWriter struct {
	Channel chan []byte
}

func (cw ChannelWriter) Write(bytes []byte) (int, error) {
	cw.Channel <- bytes
	return len(bytes), nil
}
