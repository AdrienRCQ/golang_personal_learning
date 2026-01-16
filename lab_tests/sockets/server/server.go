package server

import (
	"log"
	"net"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatalf("Couldn't set up listener : %v", err)
	}
}

func StartServer(address string) {

	listener, err := net.Listen("tcp", address)
	errorHandler(err)
	defer listener.Close()

	log.Printf("Server is ready : waiting for connection on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection missed : %v", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	log.Printf("New client connected : %s", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Connection error : %v", err)
			return
		}

		message := string(buffer[:n])
		log.Printf("Recieved : %s", message)
		conn.Write([]byte("Server says: " + message))
	}
}

func verif_message(message string) {
	command := "Aston"
	if message == command {
		log.Printf("CACA de Aston")
	} else {
		log.Printf("Toto : %s", message)
	}
}
