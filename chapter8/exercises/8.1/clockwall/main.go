// Modify clock2 to accept a port number, and write a program, clockwall, that
// acts as a client of several clock servers at once, reading the times from each one and displaying
// the results in a table, akin to the wall of clocks seen in some business offices. If you have
// access to geographically distributed computers, run instances remotely ; otherwise run local
// instances on different ports with fake time zones.
// $ TZ=US/Eastern 		./clock2 -port 8010 &
// $ TZ=Asia/Tokyo		./clock2 -port 8020 &
// $ TZ=Europe/London 	./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name string
	host string
}

func (c *clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	fmt.Println(c.name, "done")
	if s.Err() != nil {
		log.Printf("can't read from %s: %s", c.name, s.Err())
	}
}

// go run  clockwall/main.go NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall NAME=HOST ...")
		os.Exit(1)
	}

	clocks := make([]*clock, 0)
	for _, a := range os.Args[1:] {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", a)
			os.Exit(1)
		}
		clocks = append(clocks, &clock{fields[0], fields[1]})
	}

	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go c.watch(os.Stdout, conn)
	}
	// wait other goroutines do work
	for {
		time.Sleep(time.Minute)
	}
}
