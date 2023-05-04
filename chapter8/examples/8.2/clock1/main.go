package main

import (
	"io"
	"log"
	"net"
	"time"
)

// Clock1 is a TCP server that perfiodically writes the time.

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // like connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // like client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
