package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("host", "localhost", "Host of the server")
	port = flag.Int("port", 3090, "Port of the server")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	message := make(chan string)
	go MessageWrite(conn, message)
	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome, %s!", clientName)

	// Notify that a new client has arrived
	messages <- fmt.Sprintf("%s has arrived", clientName)
	incomingClients <- message // list the client channel

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s", clientName, inputMessage.Text())
	}

	// Notify that a client has left
	leavingClients <- message
	messages <- fmt.Sprintf("%s say goodbye", clientName)

}

func MessageWrite(conn net.Conn, msgs <-chan string) {
	for msg := range msgs {
		fmt.Fprintln(conn, msg)
	}
}

func Broadcaster() {
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-messages:
			for client := range clients {
				client <- msg
			}
		case clientIncoming := <-incomingClients:
			clients[clientIncoming] = true
		case clientLeaving := <-leavingClients:
			delete(clients, clientLeaving)
			close(clientLeaving)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Initiating server.")
	fmt.Println("Server listening.")

	go Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn)
	}
}
