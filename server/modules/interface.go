package main

type Module interface {
	update()
	timedUpdate()
	getId() string
}
