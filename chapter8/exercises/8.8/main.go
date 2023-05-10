// Using a select statement, add a timeout to the echo server from Section 8.3 so
// that it disconnects any client that shouts nothing within 10 seconds.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	wg := sync.WaitGroup{}
	defer func() {
		wg.Wait()
		c.Close() // ignoring errors
	}()
	timeout := 10 * time.Second
	timer := time.NewTicker(timeout)

	inputs := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			inputs <- input.Text()
		}
		if input.Err() != nil {
			log.Println("scan:", input.Err())
		}
	}()

	for {
		select {
		case input := <-inputs:
			timer.Reset(timeout)
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(c, input, 1*time.Second)
			}()
		case <-timer.C:
			fmt.Println("Closing the connection")
			return
		}
	}
}

/*
go run reverb1/main.go &
go run netcat3/main.go
*/

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
		go handleConn(conn) // handle one connection at a time
	}
}
