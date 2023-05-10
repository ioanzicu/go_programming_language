// Change the chat server’s network protocol so that each client
// provides its name on entering. Use that name instead of the network
// address when prefixing each message with its sender’s identity.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

const timeout = 5 * time.Minute

type client struct {
	Out  chan<- string
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels
			for cli := range clients {
				cli.Out <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.Out <- "Current Presents:"
			for c := range clients {
				cli.Out <- c.Name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	// who := conn.RemoteAddr().String()
	fmt.Fprint(conn, "Enter name: ")
	var who string
	_, err := fmt.Fscanln(conn, &who)
	if err != nil {
		log.Fatal(err)
		return
	}

	cli := client{ch, who}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}
	// Note: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // Note: ignoring network errors
	}
}
