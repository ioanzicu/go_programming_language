package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Clock1 is a TCP server that perfiodically writes the time.

var port = flag.Int("port", 8000, "listen port")

// go run clock/main.go -port 8010 &
// go run clock/main.go -port 8020 &
// go run clock/main.go -port 8030 &
func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // like connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
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
