package main

import (
	"fmt"
	"os"

	"github.com/etimo/go-magic-mirror/server"
)

func main() {
	bindAddress := "localhost:8080"
	fmt.Printf("Starting MAGIC-MIRROR-backend server on %s...\n", bindAddress)
	exec, _ := os.Getwd()
	fmt.Printf("Exec is: %s", exec)
	server.StartServer(bindAddress)
}
