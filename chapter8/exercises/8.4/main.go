// Modify the reverb2 server to use a sync.WaitGroup per connection to count
// the number of active echo goroutines. When it falls to zero, close the write half of the TCP
// connection as described in Exercise 8.3. Verify that your modified netcat3 client from that
// exercise waits for the final echoes of multiple concurrent shouts, even after the standard input
// has been closed.
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
	var wg sync.WaitGroup

	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		// worker

		go func() {
			echo(c, input.Text(), 1*time.Second)
			wg.Done()
		}()
	}
	wg.Wait()

	c.Close() // ignoring errors
}

/*

go run main.go &
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
