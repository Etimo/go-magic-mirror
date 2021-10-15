package models

type ChannelWriter struct {
	Channel chan []byte
}

func (cw ChannelWriter) Write(bytes []byte) (int, error) {
	copyDest := make([]byte, len(bytes))
	copy(copyDest, bytes)
	cw.Channel <- copyDest
	return len(bytes), nil
}
