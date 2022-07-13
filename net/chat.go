package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incommingClients = make(chan Client)
	leavingClients   = make(chan Client)
	messages         = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

//client1 -> server -> HandleConnection(client1)
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn, message)

	//client1:2560 Platzi.com, 38
	//platzi.com:38
	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)
	messages <- fmt.Sprintf("New client is here, name %s\n", clientName)
	incommingClients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}
	//avandonado el chat
	leavingClients <- message
	messages <- fmt.Sprintf("%s said goodbye!", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprint(conn, message)
	}
}
