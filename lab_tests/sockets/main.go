package main

import (
	"http-go/sockets/client"
	"http-go/sockets/server"
	"time"
)

func main() {
	address := "localhost"
	application_port := "8080"
	full_address := address + ":" + application_port
	go func() {
		server.StartServer(full_address)
	}()

	// Give server a moment to start
	time.Sleep(time.Second)

	// Launch client
	client.StartClient(full_address)
}
