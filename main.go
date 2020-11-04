package main

import (
	"fmt"

	"go-magic-mirror/server"
)

const bindAddress = "0.0.0.0:8080"

func main() {
	fmt.Printf("Starting server on address %s...\n", bindAddress)
	server.StartServer(bindAddress)
}
