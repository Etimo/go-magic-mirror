package modules

type Module interface {
	Update()
	TimedUpdate()
	GetId() string
	CreateFromMessage([]byte, chan []byte) (Module, error)
}
