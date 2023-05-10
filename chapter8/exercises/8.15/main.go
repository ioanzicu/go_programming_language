// Failure of any client program to read data in a timely manner
// ultimately causes all clients to get stuck. Modify the broadcaster
// to skip a message rather than wait if a client writer is not ready
// to accept it. Alternatively, add buffering to each client’s outgoing
// message channel so that most messages are not dropped; the broadcaster
// should use a non-blocking send to this channel.

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
				select {
				case cli.Out <- msg:
				default: // non-blocking receive operation
					// do nothing
				}
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
	out := make(chan string, 20) // outgoing client messages
	go clientWriter(conn, out)

	in := make(chan string) // incoming client messages
	go clientReader(conn, in)

	// who := conn.RemoteAddr().String()
	var who string
	timer := time.NewTimer(timeout)
	out <- "Enter your name:"
	select {
	case name := <-in:
		who = name
		timer.Reset(timeout)
	case <-timer.C:
		return
	}

	cli := client{out, who}
	out <- "You are " + who
	messages <- who + " is now online"
	entering <- cli

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

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}
