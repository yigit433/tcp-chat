package main

import (
	"bufio"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "3333"
	TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("Listening on %s:%s", HOST, PORT)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		msg, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Message rechieved: %s", msg)
	}
}
