package modules

type Module interface {
	Update()
	TimedUpdate()
	GetId() string
}
