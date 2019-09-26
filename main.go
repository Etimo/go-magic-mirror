package main

import (
	"fmt"

	"github.com/etimo/go-magic-mirror/server"
)

func main() {
	fmt.Printf("Starting server...\n")
	server.StartServer()
}
